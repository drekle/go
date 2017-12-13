package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/drekle/go/prometheus/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr          = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")
	observeDocker = flag.Bool("docker", true, "Observe docker statistics")
	observeK8s    = flag.Bool("kubernetes", true, "Observe kubernetes statistics")
	observeNode   = flag.Bool("node", true, "Observe node statistics")
	interval      = flag.Int("interval", 60, "Observation Interval")
)

func main() {
	flag.Parse()

	observer := metrics.NewObserver(metrics.ObserverOpts{
		Docker:     *observeDocker,
		Kubernetes: *observeK8s,
		Node:       *observeNode,
		Interval:   *interval,
	})
	observer.StartObserve()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
