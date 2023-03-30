package pingpong

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// Pong Metrics
	PongCalled = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_server_pong_called",
		Help: "The total number of times PingPong pong service was called",
	})
	PongClientCalled = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_client_pong_called",
		Help: "The total number of times PingPong client was called",
	})
	PongClientErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_client_pong_errors",
		Help: "The total number of errors PingPong pong client encountered",
	})
	PongClientLastDurationSeconds = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_client_pong_last_duration_milliseconds",
		Help: "The last observed roundtrip duration of a client call to the PingPong pong service",
	})

	// Stream Metrics
	PongStreamed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_server_stream_called",
		Help: "The total number of times PingPong stream service was called",
	})
	PongStreamedErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_server_stream_errors",
		Help: "The total number of errors PingPong stream service encountered",
	})
	PongStreamClientCalled = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_client_pong_stream_called",
		Help: "The total number of times PingPong stream client was called",
	})
	PongStreamClientErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_client_stream_errors",
		Help: "The total number of errors PingPong stream client encountered",
	})
	PongStreamClientLastDurationSeconds = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_client_pong_stream_last_duration_milliseconds",
		Help: "The last observed roundtrip duration of a client call of the PingPong stream service",
	})

	// Writer Metrics
	PongWriter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_server_writer_called",
		Help: "The total number of times PingPong writer service was called",
	})
	PongWriterErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_writer_stream_errors",
		Help: "The total number of errors PingPong writer service encountered",
	})
	WriterBytesRead = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_server_writer_last_bytes_read",
		Help: "The number of bytes last read from a file by the PingPong writer service",
	})
	WriterBytesWritten = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_server_writer_last_bytes_written",
		Help: "The number of bytes last written to a file by the PingPong writer service",
	})
	WriterClientCalled = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_client_pong_writer_called",
		Help: "The total number of times PingPong writer client was called",
	})
	PongWriterClientErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_client_writer_errors",
		Help: "The total number of errors PingPong writer client encountered",
	})
	WriterClientLastDurationSeconds = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_client_pong_writer_last_duration_milliseconds",
		Help: "The last observed roundtrip duration of a client call to the PingPong writer service",
	})

	// Proxy Metrics
	PongProxy = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_proxy_pong_called",
		Help: "The total number of times PingPong pong proxy was called",
	})
	PongProxyErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_proxy_pong_errors",
		Help: "The total number of errors PingPong pong proxy encountered",
	})
	StreamProxy = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_proxy_stream_called",
		Help: "The total number of times PingPong stream proxy was called",
	})
	StreamProxyErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_proxy_stream_errors",
		Help: "The total number of errors PingPong stream proxy encountered",
	})
	WriterProxy = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kevin_proxy_writer_called",
		Help: "The total number of times PingPong writer proxy was called",
	})
	WriterProxyErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "kevin_proxy_writer_errors",
		Help: "The total number of errors PingPong writer proxy encountered",
	})
)

func Metrics(port string, enabled bool) {
	if !enabled {
		return
	}
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		return
	}
}
