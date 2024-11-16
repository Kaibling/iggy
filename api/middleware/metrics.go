package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalRequests = prometheus.NewCounterVec( //nolint: gochecknoglobals
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var responseStatus = prometheus.NewCounterVec( //nolint: gochecknoglobals
	prometheus.CounterOpts{
		Name: "response_status_total",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{ //nolint: gochecknoglobals
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func InitMetrics() error {
	if err := prometheus.Register(totalRequests); err != nil {
		return err
	}

	if err := prometheus.Register(responseStatus); err != nil {
		return err
	}

	// todo throws error if error is checked
	prometheus.Register(httpDuration) //nolint: errcheck

	return nil
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var path string

		if strings.HasPrefix(r.URL.Path, "/api/v1") {
			path = strings.Join(strings.Split(r.URL.Path, "/")[:4], "/")
		} else {
			path = r.URL.Path
		}

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		totalRequests.WithLabelValues(path).Inc()
		responseStatus.WithLabelValues(strconv.Itoa(ww.Status())).Inc()
		timer.ObserveDuration()
	})
}
