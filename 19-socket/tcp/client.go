package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	//1. 与服务器建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//2. 利用该链接进行数据的发送和接收
	defer conn.Close()
	input := bufio.NewReader(os.Stdin) //从标准输入中获取数据的对象
	for {
		s, _ := input.ReadString('\n') // 读一行就停止
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		fmt.Println("收到服务端回复:", string(buf[:n]))
	}
}
