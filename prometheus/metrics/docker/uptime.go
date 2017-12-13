package docker

import (
	"os"

	"github.com/docker/docker/api/types"
	"github.com/drekle/go/prometheus/metrics/api"
	"github.com/prometheus/client_golang/prometheus"
)

type uptime struct {
	container types.Container
	gauge     prometheus.Gauge
}

func (metric uptime) Tick() {
	if metric.container.State == "running" {
		metric.gauge.Set(1)
	} else {
		metric.gauge.Set(0)
	}
}

func (metric uptime) GetMetric() api.Metric {
	return metric.gauge
}

func (metric uptime) GetUUID() string {
	hostname, _ := os.Hostname()
	id := hostname + "_" + metric.container.ID
	return id
}

func UptimeMetric(container types.Container) api.TickMetric {
	hostname, _ := os.Hostname()
	name := hostname + "_" + container.Names[0][1:]
	return &uptime{
		container: container,
		gauge: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: name,
				Help: "Uptime for " + name,
			},
		),
	}
}
