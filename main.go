package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	unixtimeGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "last_posted_unixtime_seconds",
			Help: "The last posted Unix time in seconds",
		},
	)
)

func init() {
	prometheus.MustRegister(unixtimeGauge)
}

func postTimestamp(w http.ResponseWriter, r *http.Request) {
	unixtimeGauge.SetToCurrentTime()
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/timestamp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postTimestamp(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

    http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
