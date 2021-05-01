package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ACCOUNT_ID                    string
	NUMBER_OF_METRICS             int
	NEW_METRIC_VALUE_FREQ_SECONDS int32
)

const PORT = ":8080"

func main() {
	ACCOUNT_ID = os.Getenv("ACC_ID")
	NEW_METRIC_VALUE_FREQ_SECONDS = 2
	NUMBER_OF_METRICS = getNumberOfMetrics()
	printEnv()
	pumpMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Printf("Server is Starting at %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
func printEnv() {
	log.Println()
	log.Printf("Label Account ID : %v ", ACCOUNT_ID)
	log.Printf("NUMBER OF COUNTER METRICS : %v ", NUMBER_OF_METRICS)
	log.Printf("NUMBER OF GAUGE METRICS : %v ", NUMBER_OF_METRICS)
	log.Printf("NEW_METRIC_VALUE_FREQ_SECONDS : %v ", NEW_METRIC_VALUE_FREQ_SECONDS)
	log.Println()

}

func getNumberOfMetrics() int {
	if s, err := strconv.Atoi(os.Args[1]); err == nil {
		fmt.Printf("Number of Metrics:  %T, %v", s, s)
		return s
	} else {
		fmt.Printf("Number of Metrics:  %T, %v", s, s)
	}

	return 0

}

func pumpMetrics() {
	for i := 1; i <= NUMBER_OF_METRICS; i++ {
		startCounter("sample_metric_counter"+strconv.Itoa(i),
			"sample metric "+strconv.Itoa(i))
		startGauge("sample_metric_gauge"+strconv.Itoa(i),
			"sample metric "+strconv.Itoa(i))
	}
}

func startCounter(Name, Help string) {
	metric := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: Name,
		Help: Help,
	}, []string{"id"})

	go func() {
		for {
			metric.WithLabelValues(ACCOUNT_ID).Inc()
			time.Sleep(time.Duration(NEW_METRIC_VALUE_FREQ_SECONDS) * time.Second)
		}
	}()
}

func startGauge(Name, Help string) {
	metric := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: Name,
		Help: Help,
	}, []string{"id"})

	go func() {
		for {
			metric.WithLabelValues(ACCOUNT_ID).Set(rand.Float64())
			time.Sleep(time.Duration(NEW_METRIC_VALUE_FREQ_SECONDS) * time.Second)
		}
	}()
}
