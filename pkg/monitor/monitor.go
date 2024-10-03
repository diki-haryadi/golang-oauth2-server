package monitor

import (
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"golang-oauth2-server/pkg/env"
)

func init() {
	registerHistogram()
	registerCounter()
}

var httpLatencyHistogram *prometheus.HistogramVec
var httpResponsesTotalCounter *prometheus.CounterVec

func registerHistogram() {
	httpLatencyHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_duration_seconds",
		Help: "the latency of http calls",
	}, []string{"handler", "method", "httpcode", "env"})

	err := prometheus.Register(httpLatencyHistogram)
	if err != nil {
		log.Printf("[Monitor] Unable to Register httpLatencyHistogram. Err: %+v\n", err)
	}
}

func registerCounter() {
	httpResponsesTotalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_responses_total",
			Help: "The count of http responses issued",
		},
		[]string{"handler", "method", "httpcode", "env"},
	)
	err := prometheus.Register(httpResponsesTotalCounter)
	if err != nil {
		log.Printf("[Monitor] Unable to Register httpResponsesTotalCounter. Err: %+v\n", err)
	}
}

// FeedHTTPMetrics to monitor latency, http code counts
func FeedHTTPMetrics(status int, duration time.Duration, path string, method string) {
	httpLatencyHistogram.With(prometheus.Labels{"handler": "all", "method": method, "httpcode": fmt.Sprintf("%d", status), "env": env.Get()}).Observe(duration.Seconds() * 1000)
	httpLatencyHistogram.With(prometheus.Labels{"handler": path, "method": method, "httpcode": fmt.Sprintf("%d", status), "env": env.Get()}).Observe(duration.Seconds() * 1000)

	httpResponsesTotalCounter.With(prometheus.Labels{"handler": "all", "method": method, "httpcode": fmt.Sprintf("%d", status), "env": env.Get()}).Inc()

}
