package main //和 main.go 在同一个文件夹下面就属于 main 包，main.go 不需要调用 main 包的文件

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

/*
server 端的基本构建：
*/

// server 的结构体封装
type Server struct {
	IP        string
	Port      int
	OnlineMap map[string]*User //在线用户的列表，key 是用户名， val 是用户指针
	mapLock   sync.RWMutex     //因为 OnlineMap 是全局的，所以需要锁，这里使用读写锁
	Message   chan string      //消息广播的 channel
}

// 创建一个 Server 的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	//fmt.Println("链接建立成功")

	// 用户上线
	user := NewUser(conn, this) //先创建 user
	user.Online()

	// 监听用户是否活跃的的 channel
	isLive := make(chan bool)

	// 接受客户端发送的消息并广播
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf) // 把客户端发送的消息读到 buf 中，返回长度 和 err
			if n == 0 {              // 说明客户端关闭了
				user.Offline()
				return
			}

			// err!=nil 说明出错了，err!=io.EOF 说明没读到最后
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// 提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			// 用户针对 msg 进行处理
			user.DoMessage(msg)

			isLive <- true // 用户的任意消息代表当前用户是一个活跃的
		}
	}()

	// 当前 Handler goroutine 先阻塞
	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，应该重置定时器
			// 但是这个 case 逻辑是不需要处理的，因为指向完之后会去判断下一个 case，判断的时候会执行 time.After 这个语句，就是更新定时器了
		case <-time.After(time.Second * 10): //time.After 是一个定时器，本质是一个管道，如果这个管道可读，说明超时了
			// 超时就关闭链接
			user.SendMsg("你因长时间未发送消息被剔除了")
			close(user.C) //  关闭管道
			conn.Close()
			return // 退出当前的 Handler
		}
	}
}

// 广播消息的方法  参数：由哪个用户发起的，发起是消息是什么
func (this *Server) BoardCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg //先拼接消息
	this.Message <- sendMsg                                  //消息放到广播的管道中
}

// 监听 Message广播消息 channel 的 goroutine，一旦有消息就发送给全部的在线 User
func (this *Server) ListenMessage() {
	for { //不断的循环
		msg := <-this.Message //先获取 广播Message 中的消息
		//将 msg 发送给所有的 user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg //把 msg 写到每个 user 的管道中
		}
		this.mapLock.Unlock()
	}
}

// 启动服务器的接口。是 Server 的成员函数
func (this *Server) Start() {
	// 1. socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port)) //拼接成addr：127.0.0.1:8888
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	defer listener.Close() //defer 实现 close

	// 启动监听 Message 的 goroutine
	go this.ListenMessage()

	for {
		// 2. accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue //如果失败 for 循环就 continue
		}

		// 3. do handler（业务处理）
		go this.Handler(conn)
	}

	// 4. close listen socket   上述的 defer 语句实现了 close
}
