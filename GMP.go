package main

import (
	"sync"
)

//系统启动时，初始化空闲P列表，并创建P,然后为了更好的工作新创建的已经准备好的协成会换线一个P,这个P会通过相关联的OS线程创建一个M,
func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		println("hello")
		wg.Done()
	}()
	go func() {

		println("world")
		wg.Done()
	}()
	wg.Wait() //等所有协成全部走完在执行
}
