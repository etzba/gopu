package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var httpRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name:      "http_request_count",
		Namespace: "gopu",
		Help:      "Number of request handled by handler",
	},
)

var numberOfConcurrentUsers = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name:      "concurrent_users",
		Namespace: "gopu",
		Help:      "Number of clients using the endpoint at the same time",
	},
)

var (
	httpRequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:      "http_request_duration",
		Namespace: "gopu",
		Help:      "Time of request processing in http handler",
	})
)
