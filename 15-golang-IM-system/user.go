package main

import (
	"net"
	"strings"
)

/*
user 类
*/

// User 中需要包含 用户名、地址、通信的管道、可以和客户端通信的链接、当前用户属于哪个 server
type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的 API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() // 获取客户端的地址
	user := &User{
		Name:   userAddr, //user 的用户名默认为 地址
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	//启动监听当前 user channel 消息的 goroutine
	go user.ListenMessage()

	return user
}

// 监听 广播Message 发送给 User channel 的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		// 不断循环 从 channel 中获取 msg 并写回 客户端
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线的业务
func (this *User) Online() {
	// 用户上线，将用户加入 onlinemap 中
	this.server.mapLock.Lock() //先加锁再往全局 map 中加数据
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BoardCast(this, "已上线")
}

// 用户下线的业务
func (this *User) Offline() {
	// 用户下线，将用户从 onlinemap 中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户下线消息
	this.server.BoardCast(this, "下线")
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" { // msg 是 who，认为用户想查询有哪些用户在线
		this.server.mapLock.Lock()
		for _, cil := range this.server.OnlineMap {
			onlineMsg := "[" + cil.Addr + "]" + cil.Name + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // 重命名的功能，如果客户端发过来的是"rename|张三"类似的消息，则认为是改名
		//name := msg[7:]
		name := strings.Split(msg, "|")[1] // 按照"|" 进行截取，截取之后返回一个字符串数组，第0个是前半段，第1个是后半段
		// 判断 name 是否被其他 user 占用了
		_, ok := this.server.OnlineMap[name]
		if ok {
			this.SendMsg("当前用户名已经被占用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[name] = this
			this.server.mapLock.Unlock()

			this.Name = name
			this.SendMsg("修改成功\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" { // 私聊模式，消息格式是"to|张三|消息内容"
		// 1 获取对方的用户名
		friendName := strings.Split(msg, "|")[1]
		if friendName == "" {
			this.SendMsg("消息格式不正确，请使用 \"to|张三|你好哦啊\"格式\n")
			return
		}
		// 2 根据用户名，得到对方 User 对象
		friendUser, ok := this.server.OnlineMap[friendName]
		if ok {
			// 3 获取消息内容，通过对方的 User 对象将消息内容发送过去
			friendMsg := strings.Split(msg, "|")[2]
			friendUser.SendMsg(this.Name + "对您说:" + friendMsg)
		} else {
			this.SendMsg("没有该好友")
			return
		}
	} else { // 否则就广播用户的消息
		this.server.BoardCast(this, msg)
	}
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
