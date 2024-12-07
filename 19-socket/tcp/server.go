package main

import (
	"fmt"
	"net"
)

// tpc server

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept() // 一致阻塞，直到链接到来
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理链接
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				buf := make([]byte, 1024)
				_, err := conn.Read(buf)
				if err != nil {
					fmt.Println("err:", err)
					break
				}
				msg := string(buf)
				fmt.Println("接收到的数据：", msg)
				conn.Write([]byte("ok"))
			}
		}(conn)
	}
}
