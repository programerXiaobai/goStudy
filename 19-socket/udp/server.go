package main

import (
	"fmt"
	"net"
)

// udp server

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer listen.Close()

	for { // udp 不再需要 accept 了
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:]) // 因为udp没有全连接队列，所以会返回 addr 是客户端的地址，我们需要给客户端的地址再返回响应
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("接收到的数据：", string(buf[:n]))

		_, err = listen.WriteToUDP([]byte("ok"), addr)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
}
