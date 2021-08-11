package consumerAndProvider

import (
	"fmt"
	"time"
)

func Produce(ch chan int)  {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println("Produce:",i)
		ch<-i
	}
}
func Consume(pipe chan int)  {
	for i := 0; i < 100; i++ {
		i2 := <-pipe
		fmt.Println("Consume:",i2)
	}
}

func Run() {
	channel := make(chan int ,3)
	go Produce(channel)

	go Consume(channel)
	time.Sleep(10*time.Second)
}