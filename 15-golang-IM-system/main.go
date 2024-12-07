package main

/*
当前进程的主入口
*/

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
