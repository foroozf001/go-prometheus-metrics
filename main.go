package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var customCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "custom_counter_total",
	Help: "The total number custom events",
})

// increment every 2 seconds
func recordMetrics() {
	go func() {
		for {
			customCounter.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
