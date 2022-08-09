package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhello1)
	http.HandleFunc("/abc", sayhello12)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func sayhello1(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world111!")
}
func sayhello12(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world1112!")
}
