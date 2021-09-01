package main

import (
	"log"
	"time"
)

//上边给你一个时间，在固定时间区域开启一个server端口
//todo 比如在进程启动10分钟内，分开启一个server端口43333，这段时间内server可以被访问到
func main() {
	timeout(5)

}

//min代表分钟  ，计时功能    ,休眠
func timeout(min int64) {
	unix := time.Now().Unix()
	log.Println(unix)
	time.Sleep(time.Duration(min) * time.Second)
	unix1 := time.Now().Unix()

	log.Println(unix1)
}

//开启一个服务端
func openServer(port int64) {

}
