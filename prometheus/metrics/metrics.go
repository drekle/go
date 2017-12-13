package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metric interface {
	prometheus.Collector
	prometheus.Metric
}

type TickMetric interface {
	Metric
	Tick()
}
