package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

/*
客户端
*/

type Client struct {
	ServerIP   string   // 连接的服务器IP
	ServerPort int      // 连接的服务器端口
	Name       string   // 连接 user 的 Name
	conn       net.Conn //连接的 socket
	flag       int      // 客户端输入的菜单选择
}

// 创建新客户端的 方法
func NewClient(serverIP string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flag:       -1, //flag 的默认值不能是0，否则 Run 函数直接就退出了
	}

	// 连接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort)) // 客户端通过 net.Dial 进行连接
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn

	return client
}

// 显示菜单的方法
func (this *Client) menu() bool {
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&this.flag) // 通过 fmt.Scanln 获取键盘中用户的选择

	if this.flag >= 0 && this.flag <= 3 {
		return true
	} else {
		fmt.Println(">>>请输入合法范围内的数字<<<")
		return false
	}
}

// 修改用户名
func (this *Client) UpdateName() bool {
	fmt.Println(">>>请输入用户名")
	fmt.Scanln(&this.Name)
	msg := "rename|" + this.Name + "\n"
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return false
	}

	return true
}

// 公聊模式
func (this *Client) PublicChat() {
	// 提示用户输入消息
	fmt.Println(">>>请输入公聊的消息，输入 exit 表示退出")
	var msg string
	fmt.Scanln(&msg)

	// for 循环，如果用户不退出，就一直处于 公聊模式
	for msg != "exit" {
		if len(msg) != 0 {
			// 发送给服务器
			_, err := this.conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("conn.Write error:", err)
			}
		}

		fmt.Println(">>>请输入公聊的消息，输入 exit 表示退出")
		msg = "" //消息制空
		fmt.Scanln(&msg)
	}
}

// 查询哪些用户在线
func (this *Client) SelectUsers() {
	msg := "who\n"
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write error:", err)
	}
}

// 私聊模式
func (this *Client) PrivateChat() {
	fmt.Println(">>>当前在线的用户如下，请输入要私聊的用户名，exit 表示退出")
	this.SelectUsers()
	var name string
	fmt.Scanln(&name)

	for name != "exit" {
		fmt.Println(">>>请输入要发送的消息,exit 表示退出")
		var msg string
		fmt.Scanln(&msg)

		for msg != "exit" {
			if len(msg) != 0 {
				privateMsg := "to|" + name + "|" + msg + "\n\n"
				_, err := this.conn.Write([]byte(privateMsg))
				if err != nil {
					fmt.Println("conn.Write error:", err)
				}
			}

			msg = ""
			fmt.Println(">>>请输入要发送的消息,exit 表示退出")
			fmt.Scanln(&msg)
		}

		name = ""
		fmt.Println(">>>当前在线的用户如下，请输入要私聊的用户名，exit 表示退出")
		this.SelectUsers()
		fmt.Scanln(&name)
	}
}

// 处理 server 回应的消息，直接显示到标准输出即可。需要是一个单独的 goroutine，否则读取 服务器回应的数据就会阻塞，不能及时处理其他的业务逻辑
func (this *Client) DealResponse() {
	io.Copy(os.Stdout, this.conn) //不断的阻塞等待 this.conn 中的数据并拷贝到 标准输出stdout，他是永久阻塞监听的。等价于下面的代码：
	//for {
	//	buf := make([]byte, 1024)
	//	_, err := this.conn.Read(buf)
	//	if err != nil {
	//		fmt.Println("conn.Read error:", err)
	//	} else {
	//		fmt.Println(string(buf))
	//	}
	//}
}

// 客户端跑的主业务
func (this *Client) Run() {
	for this.flag != 0 {
		for this.menu() == false {
		} // 只要用户输入的不合法，就一直循环让用户输入

		// 根据不同的模式处理不同的业务
		switch this.flag {
		case 1: //公聊模式

			//fmt.Println("公聊模式")
			this.PublicChat()
			//break // go 中的switch 是默认自带 break 的
		case 2: //私聊模式

			//fmt.Println("私聊模式")
			this.PrivateChat()

		case 3: //更新用户名
			//fmt.Println("更新用户名")
			this.UpdateName()
		}
	}
}

// 全局的IP 和 port 用来初始化
var serverIp string
var serverPort int

// init函数 在main 函数之前执行，用来初始化 IP 和 Port
func init() {
	// 用 go 中自带的 flag 库从命令行中读取参数，命令行格式： ./client -ip 127.0.0.1 -port 8888
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)") //第一个到最后一个的参数是:命令行输入参数的绑定变量，命令行-后面的标识，默认值，说明
	flag.IntVar(&serverPort, "port", 8888, "设置服务器port(默认是8888)")
}

// 启动客户端的 main 函数
func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>> 连接服务器失败")
		return
	}

	// 单独开启一个 goroutine 处理 server 的回执消息，防止读阻塞
	go client.DealResponse()

	fmt.Println(">>>>>>>>> 连接服务器成功")

	// 启动客户端的业务
	select {}
}
