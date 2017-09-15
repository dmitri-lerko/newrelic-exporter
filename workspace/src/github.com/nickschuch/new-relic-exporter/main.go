package main

import (
	"log"
	"net/http"

	"github.com/alecthomas/kingpin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cliPort = kingpin.Flag("port", "Port for Prometheus to scrape metrics").Default(":9000").OverrideDefaultFromEnvar("PROMETHEUS_PORT").String()
	cliPath = kingpin.Flag("/metrics", "Path for Prometheus to scrape metrics").Default(":9000").OverrideDefaultFromEnvar("PROMETHEUS_PATH").String()
	cliApp  = kingpin.Flag("application", "The name of the New Relic application").Default("").OverrideDefaultFromEnvar("NEW_RELIC_APPLICATION").String()
	cliKey  = kingpin.Flag("api-key", "The name of the New Relic application").Default("").OverrideDefaultFromEnvar("NEW_RELIC_API_KEY").String()
)

func main() {
	kingpin.Parse()

	prometheus.MustRegister(NewExporter(*cliApp, *cliKey))
	http.Handle(*cliPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*cliPort, nil))
}
