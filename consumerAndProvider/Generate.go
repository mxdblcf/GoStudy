package consumerAndProvider

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {

	ints := make(chan int, 1)
	go func() {
		for {
			ints <- rand.Int()
		}
	}()
	return ints
}

func Run2()  {

	fmt.Println(<-GenerateIntA())
	fmt.Println(<-GenerateIntA())
}