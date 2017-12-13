package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/drekle/go/prometheus/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type uptime struct {
	container *types.Container
	gauge     metrics.Metric
}

func (metric uptime) Tick() {
	if metric.container.Status == types.Healthy {
		metric.gauge.Set(1)
	} else {
		metric.gauge.Set(0)
	}
}

func UptimeMetric(container *types.Container) metrics.TickMetric {
	return &uptime{
		container: container,
		gauge: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: container.ID,
				Help: "Uptime for " + container.ID,
			},
		),
	}
}
