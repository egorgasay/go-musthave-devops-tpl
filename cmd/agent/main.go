package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	// "golang.org/x/exp/slices"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

//var counters = []string{"PollCount"}

// var metrics = []string{"Alloc", "BuckHashSys", "Frees", "GCCPUFraction", "GCSys", "HeapAlloc",
// 	"HeapIdle", "HeapInuse", "HeapReleased", "HeapSys", "LastGC", "Lookups", "MCacheInuse", "MCacheSys",
// 	"MSpanInuse", "MSpanSys", "Mallocs", "NextGC", "NumForcedGC", "NumGC", "OtherSys", "PauseTotalNs",
// 	"PollCount", "RandomValue", "StackInuse", "StackSys", "Sys", "TotalAlloc"}

const pollInterval = 2 * time.Second
const reportInterval = 10 * time.Second

//var ticker = time.NewTicker(reportInterval) //make(chan int, 29)

func main() {
	client := resty.New()
	urls := make([]string, 29)
	startReport := time.Now()
	i := 0
	for {
		start := time.Now()
		var mem runtime.MemStats
		i = 0

		urls[i] = "http://127.0.0.1:8080/update/counter/PollCount/1"
		i++

		urls[i] = fmt.Sprintf(
			"http://127.0.0.1:8080/update/gauge/RandomValue/%d",
			rand.Intn(1000000),
		)
		i++

		runtime.ReadMemStats(&mem)
		v := reflect.ValueOf(mem)
		tof := v.Type()

		for j := 0; j < v.NumField(); j++ {
			val := 0.0
			if !v.Field(j).CanUint() && !v.Field(j).CanFloat(){
				continue
			} else if !v.Field(j).CanUint() {
				val = v.Field(j).Float()
			} else {
				val = float64(v.Field(j).Uint())
			}
			
			name := tof.Field(j).Name
			urls[i] = fmt.Sprintf(
				"http://127.0.0.1:8080/update/gauge/%s/%f",
				name, val,
			)
			i++
		}
		time.Sleep(pollInterval - time.Since(start))
		if reportInterval-time.Since(startReport) <= 0 {
			startReport = time.Now()
			doRequest(urls, client)
		}
	}

}

func doRequest(urls []string, client *resty.Client) {
	fmt.Println("Отправили!")
	for _, url := range urls {
		go client.R().Post(url)
	}
}
