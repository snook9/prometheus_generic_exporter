module github.com/snook9/prometheus_generic_exporter

go 1.16

replace github.com/snook9/prometheus_generic_exporter/models => ../models

replace github.com/snook9/prometheus_generic_exporter/controllers => ../controllers

require github.com/gorilla/mux v1.8.0
