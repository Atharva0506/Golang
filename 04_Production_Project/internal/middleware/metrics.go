package middleware

import (
	"expvar"
	"net/http"
	"strconv"
	"time"
)

// Metrics are exposed at GET /debug/vars (registered automatically by the expvar package).
// In a real service you would add a Prometheus exporter alongside this.
var (
	httpRequestsTotal   = expvar.NewMap("http_requests_total")   // keyed by "<method>_<status>"
	httpRequestDuration = expvar.NewMap("http_request_duration") // keyed by "<method>_<path>", value in ms
)

// metricsResponseWriter wraps http.ResponseWriter to capture the status code and body size.
type metricsResponseWriter struct {
	http.ResponseWriter
	statusCode int
	written    int
}

func newMetricsResponseWriter(w http.ResponseWriter) *metricsResponseWriter {
	return &metricsResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (m *metricsResponseWriter) WriteHeader(code int) {
	m.statusCode = code
	m.ResponseWriter.WriteHeader(code)
}

func (m *metricsResponseWriter) Write(b []byte) (int, error) {
	n, err := m.ResponseWriter.Write(b)
	m.written += n
	return n, err
}

// Metrics returns a middleware that records per-route request counts and latencies
// using Go's built-in expvar package. Metrics are available at:
//
//	GET /debug/vars
//
// Keys exported:
//
//	http_requests_total   — map[string]*expvar.Int, key = "<METHOD>_<STATUS_CODE>"
//	http_request_duration — map[string]*expvar.Int, key = "<METHOD>_<PATH>", value = total ms
func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		mw := newMetricsResponseWriter(w)

		next.ServeHTTP(mw, r)

		duration := time.Since(start)

		// Increment request counter keyed by HTTP method + status code.
		statusKey := r.Method + "_" + strconv.Itoa(mw.statusCode)
		httpRequestsTotal.Add(statusKey, 1)

		// Accumulate latency (in milliseconds) keyed by method + path.
		routeKey := r.Method + "_" + r.URL.Path
		httpRequestDuration.Add(routeKey, duration.Milliseconds())
	})
}
