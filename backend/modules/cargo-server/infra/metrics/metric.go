package metrics

import (
	"github.com/chaos-io/gokit/metrics"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var cfg *metrics.Config
var defaultFieldKeys = []string{"method", "access_key", "error"}

func init() {
	cfg = metrics.NewConfig("metrics")
}

func GetConfig() *metrics.Config {
	return cfg
}

func NewCount(fieldKeys ...string) *kitprometheus.Counter {
	fieldKeys = append(fieldKeys, defaultFieldKeys...)
	count := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: cfg.Department,
		Subsystem: cfg.Project,
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	return count
}

func NewLatency(fieldKeys ...string) *kitprometheus.Histogram {
	fieldKeys = append(fieldKeys, defaultFieldKeys...)
	latency := kitprometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
		Namespace: cfg.Department,
		Subsystem: cfg.Project,
		Name:      "request_latency_seconds",
		Help:      "Total duration of requests in seconds.",
	}, fieldKeys)

	return latency
}
