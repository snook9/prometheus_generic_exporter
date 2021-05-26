// Name: prometheus_generic_exporter
// Authors: Jonathan CASSAING
// Prometheus exporter for various metrics.

package main

import (
	"github.com/gorilla/mux"
	"github.com/snook9/prometheus_generic_exporter/controllers"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /metrics/ to /metrics
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/metrics").Name("Index").HandlerFunc(controllers.MetricsIndex)
	return router
}
