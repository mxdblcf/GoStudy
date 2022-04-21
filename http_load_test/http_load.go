package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counts  = flag.Int("counts", 0, "request count")
	threads = flag.Int("threads", 0, "threads")
	url     = flag.String("url", "", "test url")
)

type Result struct {
	PreTotal int
	Total    int64
	Success  int64
	Fail     int64
	Error    int64
	Index    int64
}

var result Result
var startTime = time.Now().UnixNano()
var endTime int64
var wg sync.WaitGroup

func main() {
	//命令行传参使用
	flag.Parse()
	startTime = time.Now().UnixNano()

	if *counts == 0 || *threads == 0 || *url == "" {
		flag.PrintDefaults()
		return
	}

	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go gogogo(&i, *counts)
	}
	wg.Wait()
	endTime = time.Now().UnixNano()
	info()

}

func info() {
	_time := float64(endTime-startTime) / 1e9
	count := float64(result.Total) / _time
	rate := float64(result.Success) / float64(result.Total) * 100.0
	fmt.Println("*******************************")
	fmt.Println("请求总数：  ", (*counts)*(*threads))
	fmt.Println("实际请求： ", result.Total)
	fmt.Println("success： ", result.Success)
	fmt.Println("fail： ", result.Fail)
	fmt.Println("err:  ： ", result.Error)
	fmt.Println("成功率： ", fmt.Sprintf("%.2fs", rate, "%"))

	fmt.Println("每秒请求数： ", fmt.Sprintf("%.0f", count))
}
func gogogo(cIndex *int, num int) {
	for i := 0; i < num; i++ {
		atomic.AddInt64(&result.Total, 1)
		request(*url, &i, cIndex)
	}
	defer wg.Done()
}
func request(url string, i, c *int) {

	atomic.AddInt64(&result.Index, 1)
	resp, err := http.Get(url)
	if err != nil {
		atomic.AddInt64(&result.Error, 1)
		log.Println(err, i, c, &i, &c)
	} else if resp.StatusCode != 200 {
		atomic.AddInt64(&result.Fail, 1)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(resp.Body)

	} else {
		atomic.AddInt64(&result.Success, 1)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(resp.Body)
	}

}
