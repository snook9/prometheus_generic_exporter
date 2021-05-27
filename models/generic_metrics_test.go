// Name: prometheus_generic_exporter
// Authors: Jonathan CASSAING
// Prometheus exporter for various metrics.

package models

import (
	"testing"
)

func TestGetGenericMetrics(t *testing.T) {
	genericMetrics := GetGenericMetrics()
	if nil == genericMetrics {
		t.Fatalf(`Metrics are null`)
	}
}

func TestEncodeGenericMetrics(t *testing.T) {
	var genericMetrics genericMetrics
	metrics := EncodeGenericMetrics(&genericMetrics)
	if 0 >= len(*metrics) || nil == metrics {
		t.Fatalf(`Metrics encoded are null or empty`)
	}
}
