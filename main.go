package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/jamessanford/iw_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpAddr = flag.String("http", ":6798", "listen on this address")

func main() {
	flag.Parse()

	prometheus.MustRegister(collector.NewIWCollector())

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, "iw_exporter\n")
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("listening on %v\n", *httpAddr)
	log.Fatalln(http.ListenAndServe(*httpAddr, nil))
}
