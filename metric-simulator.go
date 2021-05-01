package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	mc                = kingpin.Flag("mc", "Number of metric family").Default("1").Int()
	ac                = kingpin.Flag("ac", "account id string").Default("en-100").String()
	ACCOUNT_ID        string
	NUMBER_OF_METRICS int
)

const PORT = ":8080"
const NEW_METRIC_VALUE_FREQ_SECONDS = 2

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	ACCOUNT_ID = *ac
	NUMBER_OF_METRICS = *mc
	printEnv()
	pumpMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.FileServer(http.Dir("./public")))
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
