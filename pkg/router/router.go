package router

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"strings"
	"time"

	"golang-standards-project-layout/pkg/log"
	"golang-standards-project-layout/pkg/monitor"
	"golang-standards-project-layout/pkg/response"

	"github.com/felixge/httpsnoop"
	"github.com/julienschmidt/httprouter"
)

type MyRouter struct {
	Httprouter *httprouter.Router
	Options    *Options
}

type Options struct {
	Prefix  string
	Timeout int
}

type panicObject struct {
	err        interface{}
	stackTrace string
}

type Handle func(*http.Request) *response.JSONResponse
type httpParams string
type responseWriter string

const (
	hp httpParams     = "HTTPParams"
	rw responseWriter = "ResponseWriter"
)

var (
	HttpRouter *httprouter.Router
)

func init() {
	HttpRouter = httprouter.New()
}

func NewRouter(o *Options) *MyRouter {
	myrouter := &MyRouter{Options: o}
	myrouter.Httprouter = HttpRouter
	return myrouter
}

func WrapperHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(HttpRouter, w, r)
		if r.URL.String() != "/metrics" && !strings.Contains(r.URL.String(), "/healthz") {
			monitor.FeedHTTPMetrics(m.Code, m.Duration, r.Header.Get("routePath"), r.Method)
		}
	})
}

func GetHttpParam(ctx context.Context, name string) string {
	ps := ctx.Value("HTTPParams").(httprouter.Params)
	return ps.ByName(name)
}

func GetResponseWriter(ctx context.Context) http.ResponseWriter {
	val := ctx.Value("ResponseWriter")
	if val == nil {
		return nil
	}
	return val.(http.ResponseWriter)
}

func dumpRequest(r *http.Request) []byte {
	httpDump, err := httputil.DumpRequest(r, true)
	if err == nil {
		return httpDump
	}
	log.WithFields(log.Fields{
		"url":    r.URL,
		"method": r.Method,
		"header": fmt.Sprintf("%+v", r.Header),
		"err":    err,
	}).Debugln("[Router] Failed to dump request with body, re-attempting to dump request without body")
	//Retry without including body
	httpDump, err = httputil.DumpRequest(r, false)
	if err == nil {
		return httpDump
	}
	log.WithFields(log.Fields{
		"url":    r.URL,
		"method": r.Method,
		"header": fmt.Sprintf("%+v", r.Header),
		"err":    err,
	}).Infoln("[Router] Failed to dump request")
	return nil
}

func (mr *MyRouter) handleNow(fullPath string, handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var (
			ctx context.Context
		)

		t := time.Now()
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(mr.Options.Timeout))

		defer cancel()

		ctx = context.WithValue(ctx, hp, ps)
		ctx = context.WithValue(ctx, rw, w)

		r.Header.Set("routePath", fullPath)
		r = r.WithContext(ctx)

		respChan := make(chan *response.JSONResponse)
		recovered := make(chan panicObject)

		go func() {
			defer func() {
				if err := recover(); err != nil {
					recovered <- panicObject{
						err:        err,
						stackTrace: string(debug.Stack()),
					}
				}
			}()
			respChan <- handle(r)
		}()

		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				response.NewJSONResponse().SetError(response.ErrTimeoutError).Send(w)
			}
		case cause := <-recovered:
			httpDump := dumpRequest(r)
			log.WithFields(log.Fields{
				"path":       r.URL.Path,
				"httpDump":   string(httpDump),
				"stackTrace": cause.stackTrace,
				"error":      fmt.Sprintf("%v", cause.err),
			}).Errorln("[Router] panic have occurred")
			response.NewJSONResponse().SetError(response.ErrInternalServerError).Send(w)
		case resp := <-respChan:
			if resp != nil {
				resp.SetLatency(time.Since(t).Seconds() * 1000)
				if resp.StatusCode > 499 {
					m := map[string]interface{}{}
					httpDump := dumpRequest(r)
					m["ERROR:"] = resp.RealError
					m["RESPONSE:"] = string(resp.GetBody())
					m["DUMP:"] = string(httpDump)
					log.Printf("%+v", m)
				}
				resp.Send(w)
			} else {
				if fullPath != "/metrics" {
					httpDump := dumpRequest(r)
					log.WithFields(log.Fields{
						"dump": string(httpDump),
					}).Errorln("[Router] Nil response received from the handler")
					response.NewJSONResponse().SetError(response.ErrInternalServerError).Send(w)
				}
			}
		}
	}
}

func (mr *MyRouter) Handle(path, method string, handle Handle) {
	fullPath := mr.Options.Prefix + path
	log.Println(fullPath)
	mr.Httprouter.Handle(method, fullPath, mr.handleNow(fullPath, handle))
}

func (mr *MyRouter) GET(path string, handle Handle) {
	mr.Handle(path, http.MethodGet, handle)
}

func (mr *MyRouter) POST(path string, handle Handle) {
	mr.Handle(path, http.MethodPost, handle)
}

func (mr *MyRouter) PUT(path string, handle Handle) {
	mr.Handle(path, http.MethodPut, handle)
}

func (mr *MyRouter) PATCH(path string, handle Handle) {
	mr.Handle(path, http.MethodPatch, handle)
}

func (mr *MyRouter) DELETE(path string, handle Handle) {
	mr.Handle(path, http.MethodDelete, handle)
}

func (mr *MyRouter) ServeFiles(path string, root http.FileSystem) {
	fullPath := mr.Options.Prefix + path
	mr.Httprouter.ServeFiles(fullPath, root)
}

func (mr *MyRouter) Group(path string, fn func(r *MyRouter)) {
	sr := NewRouter(&Options{
		Prefix:  mr.Options.Prefix + path,
		Timeout: mr.Options.Timeout,
	})
	fn(sr)
}
