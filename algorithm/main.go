package main

import (
	"GoStudy/algorithm/sort"
	"fmt"
)

func main() {

	arr := [...]int{1, 9, 6, 33, 2, 45, 7, 4523}
	sort.Bsort(arr[:])
	fmt.Println(arr)
}
