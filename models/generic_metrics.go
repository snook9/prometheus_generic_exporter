// Name: prometheus_generic_exporter
// Authors: Jonathan CASSAING
// Prometheus exporter for various metrics.

package models

import (
	"fmt"
	"time"
)

type genericMetrics struct {
	Hashrate  float32
	Name      string
	CreatedAt time.Time
}

func GetGenericMetrics() *genericMetrics {
	var genericMetrics genericMetrics

	// Read all metrics here
	genericMetrics.Name = "Mes super metrics"
	genericMetrics.Hashrate = 135.40
	genericMetrics.CreatedAt = time.Now()

	return &genericMetrics
}

//Encode a structure in prometheus exporter format
func EncodeGenericMetrics(genericMetrics *genericMetrics) *string {
	metrics := `generic_hashrate{name="` + genericMetrics.Name + `"` +
		`,creation_date="` + genericMetrics.CreatedAt.String() + `"} ` +
		fmt.Sprintf("%f", genericMetrics.Hashrate)

	return &metrics
}
