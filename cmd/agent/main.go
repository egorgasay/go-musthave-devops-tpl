package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

// var metrics = []string{"Alloc", "BuckHashSys", "Frees", "GCCPUFraction", "GCSys", "HeapAlloc",
// 	"HeapIdle", "HeapInuse", "HeapReleased", "HeapSys", "LastGC", "Lookups", "MCacheInuse", "MCacheSys",
// 	"MSpanInuse", "MSpanSys", "Mallocs", "NextGC", "NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs",
// 	"PollCount", "RandomValue", "StackInuse", "StackSys", "Sys", "TotalAlloc"}

var pollInterval = 2 * time.Second
var reportInterval = 10 * time.Second
var baseURL = "localhost:8080"

//var ticker = time.NewTicker(reportInterval) //make(chan int, 29)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func main() {
	startReport := time.Now()

	if addr, ok := os.LookupEnv("ADDRESS"); ok {
		baseURL = addr
	}

	if repInterval, ok := os.LookupEnv("REPORT_INTERVAL"); ok {
		sec, err := strconv.Atoi(repInterval)
		if err == nil {
			reportInterval = time.Duration(sec) * time.Second
		}
	}

	if pollInterv, ok := os.LookupEnv("POLL_INTERVAL"); ok {
		sec, err := strconv.Atoi(pollInterv)
		if err == nil {
			pollInterval = time.Duration(sec) * time.Second
		}
	}

	fmt.Println(pollInterval, reportInterval, baseURL)

	for {
		var requests = make([]*resty.Request, 0, 29)
		start := time.Now()
		var mem runtime.MemStats

		requests = makeNewRequest("counter", "PollCount", 1.0, requests)
		requests = makeNewRequest("gauge", "RandomValue", float64(rand.Intn(1000000)), requests)

		runtime.ReadMemStats(&mem)
		v := reflect.ValueOf(mem)
		tof := v.Type()

		for j := 0; j < v.NumField(); j++ {
			val := 0.0
			if !v.Field(j).CanUint() && !v.Field(j).CanFloat() {
				continue
			} else if !v.Field(j).CanUint() {
				val = v.Field(j).Float()
			} else {
				val = float64(v.Field(j).Uint())
			}

			name := tof.Field(j).Name
			requests = makeNewRequest("gauge", name, val, requests)
		}
		time.Sleep(pollInterval - time.Since(start))
		if reportInterval-time.Since(startReport) <= 0 {
			startReport = time.Now()
			doRequest(requests)
		}
	}

}

func makeNewRequest(mtype, id string, val float64, requests []*resty.Request) []*resty.Request {
	cli := resty.New().SetBaseURL("http://" + baseURL)
	var mt Metrics
	if mtype == "gauge" {
		mt.MType = "gauge"
		mt.Value = &val
	} else if mtype == "counter" {
		mt.MType = "counter"
		delta := int64(val)
		mt.Delta = &delta
	}

	mt.ID = id
	req := cli.R().SetHeader("Content-Type", "application/json").SetBody(&mt)
	requests = append(requests, req)
	return requests
}

func doRequest(requests []*resty.Request) {
	fmt.Println("Отправили!")
	for _, req := range requests {
		go req.Post("update/")
	}
}
