package GoStudy

//思路，通过 tcp将安全内核的数据拿到，然后进行分类，传递给下一层server
import (
	"fmt"

	"log"
	"net"
)

var (
	strings = make(chan string)
)

//监听
func listener() {
	listener, err1 := net.Listen("tcp", "127.0.0.1:9001")
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)

	//2、接收多个用户连接
	for {
		conn, err2 := listener.Accept()
		if err2 != nil {
			log.Println(err2)
			return
		}
		//3、处理用户请求，新建一个协程
		go Handle(conn)
	}
}

//处理器进行数据整理
func Handle(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn) //函数调用完毕，自动关闭conn

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

		//从上游服务接来的data,发送到通道中
		strings <- string(buf[:n])

		//打印
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))

		//如果接收的数据是exit，则断开该
		if "exit" == string(buf[:n-2]) { //接收到的数据包含/r/n时，去掉后面两个字符
			fmt.Printf("%v exit\n", addr)
			return
		}
		//todo data分类整理,出来ABC三类数据，给下游客户端abc
	}

}
