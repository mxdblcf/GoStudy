package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close() //函数调用完毕，自动关闭conn

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect successful")

	buf := make([]byte, 2048)
	//循环读取接收到的数据，直到客户端断开
	for {
		n, err3 := conn.Read(buf)
		if err3 != nil {
			fmt.Println("err3 = ", err3)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		//如果接收的数据是exit，则断开该客户端
		if "exit" == string(buf[:n-2]) { //接收到的数据包含/r/n时，去掉后面两个字符
			// if "exit" == string(buf[:n]) {
			fmt.Printf("%v exit\n", addr)
			return
		}

		//把数据转换为大写，再发送给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	//1、监听
	listener, err1 := net.Listen("tcp", "127.0.0.1:9001")
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}
	defer listener.Close()

	//2、接收多个用户连接
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("err2 = ", err2)
			return
		}
		//3、处理用户请求，新建一个协程
		go HandleConn(conn)
	}
}
