package main

import (
	"sync"

	"github.com/previousnext/go-newrelic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const namespace = "newrelic"

// Exporter is a custom Prometheus metrics exporter.
type Exporter struct {
	mutex sync.Mutex

	name string
	key  string

	responseTime  *prometheus.Desc
	throughput    *prometheus.Desc
	errorRate     *prometheus.Desc
	apdexTarget   *prometheus.Desc
	apdexScore    *prometheus.Desc
	hostCount     *prometheus.Desc
	instanceCount *prometheus.Desc
}

// NewExporter instantiates a new Prometheus metrics exporter.
func NewExporter(name, key string) *Exporter {
	labels := map[string]string{
		"newrelic_application": name,
	}

	return &Exporter{
		name: name,
		key:  key,
		responseTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "response_time"),
			"The duration of time between a request for service and a response.",
			nil,
			labels),
		throughput: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "throughput"),
			"Requests per minute (RPM)",
			nil,
			labels),
		errorRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "error_rate"),
			"Rate of errors responses",
			nil,
			labels),
		apdexTarget: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "apdex_target"),
			"User specified target for Apdex score",
			nil,
			labels),
		apdexScore: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "apdex_score"),
			"Industry-standard way to measure users' satisfaction with the response time of an application or service",
			nil,
			labels),
		hostCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "host_count"),
			"Number of hosts",
			nil,
			labels),
		instanceCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "instance_count"),
			"Number of instances",
			nil,
			labels),
	}
}

// Describe is used to list our New Relic metrics.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.responseTime
	ch <- e.throughput
	ch <- e.errorRate
	ch <- e.apdexTarget
	ch <- e.apdexScore
	ch <- e.hostCount
	ch <- e.instanceCount
}

// Collect is collects our new metrics.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// To protect metrics from concurrent collects.
	e.mutex.Lock()
	defer e.mutex.Unlock()

	app, err := newrelic.New(e.key).Application(e.name)
	if err != nil {
		log.Errorf("Failed to load New Relic application: %s", err)
	}

	ch <- prometheus.MustNewConstMetric(e.responseTime, prometheus.CounterValue, app.ApplicationSummary.ResponseTime)
	ch <- prometheus.MustNewConstMetric(e.throughput, prometheus.CounterValue, app.ApplicationSummary.Throughput)
	ch <- prometheus.MustNewConstMetric(e.errorRate, prometheus.CounterValue, app.ApplicationSummary.ErrorRate)
	ch <- prometheus.MustNewConstMetric(e.apdexTarget, prometheus.CounterValue, app.ApplicationSummary.ApdexTarget)
	ch <- prometheus.MustNewConstMetric(e.apdexScore, prometheus.CounterValue, app.ApplicationSummary.ApdexScore)
	ch <- prometheus.MustNewConstMetric(e.hostCount, prometheus.CounterValue, app.ApplicationSummary.HostCount)
	ch <- prometheus.MustNewConstMetric(e.instanceCount, prometheus.CounterValue, app.ApplicationSummary.InstanceCount)
}
