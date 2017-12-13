package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
)

type DockerMetrics struct {
	ticker *time.Ticker
	cli    *client.Client
	//TODO: Support all prometheus interfaces
	metrics map[string]prometheus.Gauge
}

type DockerUptime interface {
	prometheus.Gauge
	metrics.Metric
}

func (metric DockerUptime) collect() {
	metric.cli.ContainerStats()
	metric.Set(1)
}

var metrics *DockerMetrics

func MetricsInstance() (*DockerMetrics, error) {
	if metrics == nil {
		envCli, err := client.NewEnvClient()
		if err != nil {
			return nil, err
		}
		metrics = &DockerMetrics{
			ticker:  time.NewTicker(60),
			cli:     envCli,
			metrics: make(map[string]interface{}),
		}
	}
	return metrics, nil
}

func (metrics *DockerMetrics) StartObserve() {
	go func() {
		var ok bool
		for {
			_, ok := <-metrics.ticker.C
			if !ok {
				break
			}
			containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
			if err != nil {
				panic(err)
			}

			for _, container := range containers {
				fmt.Printf("%s %s\n", container.ID[:10], container.Image)
			}
		}
	}()
}

func (metrics *DockerMetrics) StopObserve() {
	close(metrics.ticker)
}

func (metrics *DockerMetrics) RegisterAll() {
	close(metrics.ticker)
}
