package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		str, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		// 接收数据
		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[:]) // 必须传进去一个切片，要么 buf 本来就是切片；如果 buf 定义的是一个数组，需要写一个 buf[:] 变成一个切片
		if err != nil {
			fmt.Println("err:", err)
			return
		}

		fmt.Printf("read from %v, msg:%v\n", addr, string(buf[:n]))
	}
}
