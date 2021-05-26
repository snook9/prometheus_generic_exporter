// Name: prometheus_generic_exporter
// Authors: Jonathan CASSAING
// Prometheus exporter for various metrics.

package controllers

import (
	"net/http"

	"github.com/snook9/prometheus_generic_exporter/models"
)

func MetricsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	binancePoolMetrics := models.GetBinancePoolMetrics()

	buffer := []byte(*models.EncodeBinancePoolMetrics(binancePoolMetrics))

	w.Write(buffer)
}
