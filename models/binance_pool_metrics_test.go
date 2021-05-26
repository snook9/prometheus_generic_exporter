// Name: prometheus_generic_exporter
// Authors: Jonathan CASSAING
// Prometheus exporter for various metrics.

package models

import (
	"testing"
)

func TestGetBinancePoolMetrics(t *testing.T) {
	binancePoolMetrics := GetBinancePoolMetrics()
	if nil == binancePoolMetrics {
		t.Fatalf(`Metrics are null`)
	}
}

func TestEncodeBinancePoolMetrics(t *testing.T) {
	var binancePoolMetrics binancePoolMetrics
	metrics := EncodeBinancePoolMetrics(&binancePoolMetrics)
	if 0 >= len(*metrics) || nil == metrics {
		t.Fatalf(`Metrics encoded are null or empty`)
	}
}
