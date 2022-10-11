package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"io"
	"time"
)

const addr = "localhost:9999"
const message = "ccc"

func QuicCli() {
	session, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println(message)
		_, err = stream.Write([]byte(message))
		if err != nil {
			fmt.Println(err)
			return
		}
		buf := make([]byte, len(message))
		_, err = io.ReadFull(stream, buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(buf)
		time.Sleep(2 * time.Second)
	}
}
