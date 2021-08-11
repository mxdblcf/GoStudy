package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",sayhello1)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func sayhello1(resp http.ResponseWriter,req *http.Request){
	io.WriteString(resp,"Hello world111!")
}