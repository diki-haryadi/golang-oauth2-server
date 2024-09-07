package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"golang-standards-project-layout/pkg/constant"
	"golang-standards-project-layout/pkg/custerr"
	"golang-standards-project-layout/pkg/log"
)

var (
	ErrBadRequest          = errors.New("bad request")
	ErrForbiddenResource   = errors.New("forbidden resource")
	ErrNotFound            = errors.New("not found")
	ErrPreConditionFailed  = errors.New("precondition failed")
	ErrInternalServerError = errors.New("internal server error")
	ErrTimeoutError        = errors.New("timeout error")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrConflict            = errors.New("conflict")
	ErrNoContent           = errors.New("no content")
)

const (
	StatusCodeGenericSuccess            = "200000"
	StatusCodeNoContent                 = "204000"
	StatusCodeBadRequest                = "400000"
	StatusCodeAlreadyRegistered         = "400001"
	StatusCodeUnauthorized              = "401000"
	StatusCodeForbidden                 = "403000"
	StatusCodeNotFound                  = "404000"
	StatusCodeConflict                  = "409000"
	StatusCodeGenericPreconditionFailed = "412000"
	StatusCodeOTPLimitReached           = "412550"
	StatusCodeNoLinkerExist             = "412553"
	StatusCodeInternalError             = "500000"
	StatusCodeFailedSellBatch           = "500100"
	StatusCodeFailedOTP                 = "503000"
	StatusCodeServiceUnavailable        = "503000"
	StatusCodeTimeoutError              = "504000"
)

func getErrType(err error) error {
	switch e := err.(type) {
	case custerr.ErrChain:
		errType := err.(custerr.ErrChain).Type
		if errType != nil {
			err = errType
		}
	default:
		_ = e
	}
	return err
}

func GetErrorCode(err error) string {
	err = getErrType(err)

	switch err {
	case ErrBadRequest:
		return StatusCodeBadRequest
	case ErrForbiddenResource:
		return StatusCodeForbidden
	case ErrNotFound:
		return StatusCodeNotFound
	case ErrConflict:
		return StatusCodeConflict
	case ErrUnauthorized:
		return StatusCodeUnauthorized
	case ErrForbiddenResource:
		return StatusCodeForbidden
	case ErrPreConditionFailed:
		return StatusCodeGenericPreconditionFailed
	case ErrInternalServerError:
		return StatusCodeInternalError
	case ErrTimeoutError:
		return StatusCodeTimeoutError
	case nil:
		return StatusCodeGenericSuccess
	case ErrNoContent:
		return StatusCodeNoContent
	default:
		return StatusCodeInternalError
	}
}

func GetHTTPCode(code string) int {
	s := code[0:3]
	i, _ := strconv.Atoi(s)
	return i
}

type JSONResponse struct {
	Data        interface{}            `json:"data,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Code        string                 `json:"code"`
	StatusCode  int                    `json:"statusCode"`
	ErrorString string                 `json:"error,omitempty"`
	Error       error                  `json:"-"`
	RealError   string                 `json:"-"`
	Latency     string                 `json:"latency"`
	Log         map[string]interface{} `json:"-"`
	HTMLPage    bool                   `json:"-"`
	Result      interface{}            `json:"result,omitempty"`
}

func NewJSONResponse() *JSONResponse {
	return &JSONResponse{Code: StatusCodeGenericSuccess, StatusCode: GetHTTPCode(StatusCodeGenericSuccess), Log: map[string]interface{}{}}
}

func (r *JSONResponse) SetData(data interface{}) *JSONResponse {
	r.Data = data
	return r
}

func (r *JSONResponse) SetHTML() *JSONResponse {
	r.HTMLPage = true
	return r
}

func (r *JSONResponse) SetResult(result interface{}) *JSONResponse {
	r.Result = result
	return r
}

func (r *JSONResponse) SetMessage(msg string) *JSONResponse {
	r.Message = msg
	return r
}

func (r *JSONResponse) SetLatency(latency float64) *JSONResponse {
	r.Latency = fmt.Sprintf("%.2f ms", latency)
	return r
}

func (r *JSONResponse) APIStatusSuccess() *JSONResponse {
	r.Code = constant.StatusCode(constant.StatusSuccess)
	r.Message = constant.StatusText(constant.StatusSuccess)
	return r
}

func (r *JSONResponse) APIStatusCreated() *JSONResponse {
	r.StatusCode = constant.StatusCreated
	r.Code = constant.StatusCode(constant.StatusCreated)
	r.Message = constant.StatusText(constant.StatusCreated)
	return r
}

func (r *JSONResponse) APIStatusAccepted() *JSONResponse {
	r.StatusCode = constant.StatusAccepted
	r.Code = constant.StatusCode(constant.StatusAccepted)
	r.Message = constant.StatusText(constant.StatusAccepted)
	return r
}

func (r *JSONResponse) APIStatusErrorUnknown() *JSONResponse {
	r.StatusCode = constant.StatusErrorUnknown
	r.Code = constant.StatusCode(constant.StatusErrorUnknown)
	r.Message = constant.StatusText(constant.StatusErrorUnknown)
	return r
}

func (r *JSONResponse) APIStatusInvalidAuthentication() *JSONResponse {
	r.StatusCode = constant.StatusInvalidAuthentication
	r.Code = constant.StatusCode(constant.StatusInvalidAuthentication)
	r.Message = constant.StatusText(constant.StatusInvalidAuthentication)
	return r
}

func (r *JSONResponse) APIStatusUnauthorized() *JSONResponse {
	r.StatusCode = constant.StatusUnauthorized
	r.Code = constant.StatusCode(constant.StatusUnauthorized)
	r.Message = constant.StatusText(constant.StatusUnauthorized)
	return r
}

func (r *JSONResponse) APIStatusForbidden() *JSONResponse {
	r.StatusCode = constant.StatusForbidden
	r.Code = constant.StatusCode(constant.StatusForbidden)
	r.Message = constant.StatusText(constant.StatusForbidden)
	return r
}

func (r *JSONResponse) APIStatusBadRequest() *JSONResponse {
	r.StatusCode = constant.StatusErrorForm
	r.Code = constant.StatusCode(constant.StatusErrorForm)
	r.Message = constant.StatusText(constant.StatusErrorForm)
	return r
}

func (r *JSONResponse) APIStatusNotFound() *JSONResponse {
	r.StatusCode = constant.StatusNotFound
	r.Code = constant.StatusCode(constant.StatusNotFound)
	r.Message = constant.StatusText(constant.StatusNotFound)
	return r
}

func (r *JSONResponse) SetLog(key string, val interface{}) *JSONResponse {
	_, file, no, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"code":            r.Code,
		"err":             val,
		"function_caller": fmt.Sprintf("file %v line no %v", file, no),
	}).Errorln("Error API")
	r.Log[key] = val
	return r
}

func (r *JSONResponse) SetError(err error, a ...string) *JSONResponse {
	r.Code = GetErrorCode(err)
	r.SetLog("error", err)
	r.RealError = fmt.Sprintf("%+v", err)
	err = getErrType(err)
	r.Error = err
	r.ErrorString = err.Error()
	r.StatusCode = GetHTTPCode(r.Code)
	if r.Code == StatusCodeNoContent {
		r.StatusCode = GetHTTPCode(StatusCodeBadRequest)
	}
	if r.StatusCode == http.StatusInternalServerError {
		r.ErrorString = "Internal Server error"
	}
	if len(a) > 0 {
		r.ErrorString = a[0]
	}
	return r
}

func (r *JSONResponse) GetBody() []byte {
	b, _ := json.Marshal(r)
	return b
}

func (r *JSONResponse) Send(w http.ResponseWriter) {
	if r.HTMLPage {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(r.StatusCode)
	} else {
		w.Header().Set("Content-Type", "application/json")
		statusCode := r.StatusCode
		if r.Code == StatusCodeNoContent {
			statusCode = http.StatusOK
			r.Code = StatusCodeNoContent
		}
		w.WriteHeader(statusCode)
		err := json.NewEncoder(w).Encode(r)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Errorln("[JSONResponse] Error encoding response")
		}
	}
}
