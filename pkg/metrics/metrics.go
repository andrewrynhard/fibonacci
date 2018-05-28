package metrics

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// DurationHistogramName is a prometheus metric name.
	DurationHistogramName = "sequence_duration_seconds"
	// DurationHistogram is a prometheus metric.
	DurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: DurationHistogramName,
		Help: "Time taken to create sequences",
	}, []string{"code"})

	// CacheHitsCounterName is a prometheus metric name.
	CacheHitsCounterName = "cache_hits"
	// CacheHitsCounter is a prometheus metric.
	CacheHitsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: CacheHitsCounterName,
		Help: "The number of cache hits.",
	})

	// CacheMissesCounterName is a prometheus metric name.
	CacheMissesCounterName = "cache_misses"
	// CacheMissesCounter is a prometheus metric.
	CacheMissesCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: CacheMissesCounterName,
		Help: "The number of cache misses.",
	})
)

// ServeMetrics starts an HTTP server for use by Prometheus.
func ServeMetrics(port int) (err error) {
	prometheus.MustRegister(DurationHistogram)
	prometheus.MustRegister(CacheHitsCounter)
	prometheus.MustRegister(CacheMissesCounter)

	http.Handle("/metrics", promhttp.Handler())

	if err = http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		return
	}

	return nil
}
