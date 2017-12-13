package docker

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/drekle/go/prometheus/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type DockerMetrics struct {
	ticker *time.Ticker
	cli    *client.Client
	//TODO: Support all prometheus interfaces
	metrics map[string]*metrics.TickMetric
}

var instance *DockerMetrics

func MetricsInstance() (*DockerMetrics, error) {
	if instance == nil {
		envCli, err := client.NewEnvClient()
		if err != nil {
			return nil, err
		}
		instance = &DockerMetrics{
			//Todo: Configurable
			cli:     envCli,
			metrics: make(map[string]*metrics.TickMetric),
		}
	}
	return instance, nil
}

func (instance *DockerMetrics) StartObserve() {

	instance.ticker = time.NewTicker(60)
	go func() {
		var ok bool
		for {
			_, ok := <-instance.ticker.C
			if !ok {
				break
			}
			containers, err := instance.cli.ContainerList(context.Background(), types.ContainerListOptions{})
			if err != nil {
				panic(err)
			}

			for _, container := range containers {
				metric := UptimeMetric(&container)
				prometheus.MustRegister(metric)
			}
		}
	}()
}

func (instance *DockerMetrics) StopObserve() {
	instance.ticker.Stop()
}

func (metrics *DockerMetrics) RegisterAll() {
}
