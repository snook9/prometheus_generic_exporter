// Name: prometheus_generic_exporter
// Authors: Jonathan CASSAING
// Prometheus exporter for various metrics.

package models

import (
	"fmt"
	"time"
)

type binancePoolMetrics struct {
	Hashrate  float32   //`json:"hashrate"`
	Name      string    //`json:"name"`
	CreatedAt time.Time //`json:"created_at"`
}

func GetBinancePoolMetrics() *string {
	var binancePoolMetrics binancePoolMetrics

	// Read all metrics here
	binancePoolMetrics.Name = "Mon pool"
	binancePoolMetrics.Hashrate = 135.40
	binancePoolMetrics.CreatedAt = time.Now()

	metrics := "binance_pool_hashrate{name=\"" + binancePoolMetrics.Name + "\"" +
		",creation_date=\"" + binancePoolMetrics.CreatedAt.String() + "\"} " +
		fmt.Sprintf("%f", binancePoolMetrics.Hashrate)

	return &metrics
}
