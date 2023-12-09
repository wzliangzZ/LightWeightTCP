package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	// 连接服务器，得到一个conn
	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")

	if err != nil {
		fmt.Println("client err:", err)
		return
	}

	for {
		// 连接调用Write,写数据
		_, err := conn.Write([]byte("hello Zinx v0.1"))
		if err != nil {
			fmt.Println("clent write err ", err)
			return
		}
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			return
		}

		fmt.Printf("server call back:%s cnt = %d \n", buf, n)
		// 阻塞
		time.Sleep(time.Second)
	}

}
