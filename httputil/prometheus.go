package httputil

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var c = prometheus.NewCounterVec(
	prometheus.CounterOpts{Name: "api_request_total"},
	[]string{"path", "method"},
)

func init() {
	prometheus.MustRegister(c)
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("I am here."))
}

/*
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api", APIHandler)
	r.Handle("/metrics", promhttp.Handler())

	r.Use(counterMiddleware)
	_ = http.ListenAndServe(":11111", r)

}
*/

func counterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.With(prometheus.Labels{
			"path":   r.RequestURI,
			"method": r.Method,
		}).Inc()
		next.ServeHTTP(w, r)
	})
}
