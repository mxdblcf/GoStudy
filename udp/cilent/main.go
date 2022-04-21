package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	write, err := conn.Write([]byte("sdsadsa"))
	fmt.Println(write)
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	data := make([]byte, 4)
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}
	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
	os.Exit(0)
}
