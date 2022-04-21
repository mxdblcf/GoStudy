package main

import "fmt"

func main() {
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int64{1, 2, 3, 45, 6}
	for i, i2 := range slice {
		fmt.Println(i, ":", i2)
	}
}
