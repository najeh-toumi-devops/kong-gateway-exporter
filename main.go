package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "time"
    "io/ioutil"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    kongURL string
    port    string
)

var kongUp = prometheus.NewGauge(prometheus.GaugeOpts{
    Name: "kong_up",
    Help: "1 if Kong is up, 0 if down",
})

func init() {
    prometheus.MustRegister(kongUp)
}

func scrapeKong() {
    client := http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(kongURL)
    if err != nil {
        kongUp.Set(0)
        return
    }
    defer resp.Body.Close()
    _, _ = ioutil.ReadAll(resp.Body)
    if resp.StatusCode == 200 {
        kongUp.Set(1)
    } else {
        kongUp.Set(0)
    }
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    scrapeKong()
    promhttp.Handler().ServeHTTP(w, r)
}

func main() {
    flag.StringVar(&kongURL, "kong-url", "http://localhost:8001/", "Kong Admin API URL")
    flag.StringVar(&port, "port", "9542", "Port to expose metrics")
    flag.Parse()

    http.HandleFunc("/metrics", metricsHandler)
    fmt.Printf("Kong Gateway Exporter running on :%s/metrics\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
