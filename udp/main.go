package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func init() {
	fmt.Println("no1")
}
func main() {

	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8888")

	listen, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	defer listen.Close()
	for {
		handleClient(listen)
	}
}
func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	fmt.Println(string(data))
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	daytime := time.Now().Unix()
	fmt.Println(n, remoteAddr)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(daytime))
	conn.WriteToUDP(b, remoteAddr)
}
