package main

import (
	"fmt"
	"net"
)

//连续发送20条数据 服务器接收到的并不是20条记录 而是存在粘包情况 多条数据封装在同一个包内
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello,Thank you. Are you OK Indian Mi fan`
		conn.Write([]byte(msg))
	}
}
