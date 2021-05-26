// Name: prometheus_generic_exporter
// Authors: Jonathan CASSAING
// Prometheus exporter for various metrics.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var programName string = "Prometheus Generic Exporter"
var programVersion string = "0.1.0"

func main() {
	versionPtr := flag.Bool("version", false, "Display the version of the software")
	portPtr := flag.Int("listen-port", 8080, "Listen TCP port")

	flag.Parse()

	if true == *versionPtr {
		fmt.Printf("%v v%v\n", programName, programVersion)
		os.Exit(0)
	}

	fmt.Printf("%v v%v\n", programName, programVersion)
	fmt.Println("Program started!")
	fmt.Printf("Listening on port %d...\n", *portPtr)
	fmt.Printf("Try http://localhost:%d/metrics\n", *portPtr)

	router := InitializeRouter()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*portPtr), router))
}
