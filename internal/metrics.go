package internal

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Metrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
