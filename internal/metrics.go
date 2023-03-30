package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	PongCalled = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_pong_called",
		Help: "The total number of times PingPong service was called",
	})
	PongStreamed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_stream_called",
		Help: "The total number of times PingPong service stream was called",
	})
	PongWriter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_writer_called",
		Help: "The total number of times PingPong writer service was called",
	})
	WriterBytesRead = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_writer_last_bytes_read",
		Help: "The number of bytes last read from a file by the PingPong writer service",
	})
	WriterBytesWritten = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_writer_last_bytes_written",
		Help: "The number of bytes last written to a file by the PingPong writer service",
	})
)

func Metrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
