package main

import (
	"flag"
	"fmt"
)

func main() {
	surname := flag.String("surname", "王", "您的姓氏")
	var personName string
	flag.StringVar(&personName, "personName", "小二", "您的名")
	id := flag.Int("id", 0, "您的Id")
	flag.Parse()
	fmt.Printf("i am %v %v ,and my id is %v \n", *surname, personName, *id)

}

//flag命令读取参数    go run test1.go   -surname="m" -personName="x" -id=100
