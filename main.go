package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)





func main() {
	fmt.Println("AWS MFA Exporter")

	metricsExample := prometheus.NewDesc(
		prometheus.BuildFQName("namespace", "", "messages_received_total"),
		"How many messages have been received (per channel).",
		[]string{"channel"}, nil,
	   )

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9101", nil))

	fmt.Println(metricsExample)
}
