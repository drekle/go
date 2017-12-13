package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metric interface {
	Tick()
	prometheus.Collector
}
