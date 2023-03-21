package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	PongCalled = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_pong_called",
		Help: "The total number of times PingPong service was called",
	})
	PongStreamed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_stream_called",
		Help: "The total number of times PingPong service stream was called",
	})
)

func Metrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
