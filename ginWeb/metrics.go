package main

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	RequestDuration *prometheus.HistogramVec
	RequestTotal    *prometheus.CounterVec
	ErrorsTotal     *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_server_request_duration_seconds",
				Help:    "Duration of HTTP requests",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"pod_id", "datacenter", "method", "route"},
		),

		RequestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_server_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"pod_id", "datacenter", "method", "route", "status_code"},
		),

		ErrorsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_server_errors_total",
				Help: "Total number of HTTP 4xx and 5xx",
			},
			[]string{"pod_id", "datacenter", "method", "route", "status_code"},
		),
	}
}

func (m *Metrics) Register(reg prometheus.Registerer) error {
	if err := reg.Register(m.RequestDuration); err != nil {
		return err
	}
	if err := reg.Register(m.RequestTotal); err != nil {
		return err
	}
	if err := reg.Register(m.ErrorsTotal); err != nil {
		return err
	}

	return nil
}
