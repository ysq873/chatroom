package main

import (
	"fmt"
	"net"
)

// 处理登录信息
func process(conn net.Conn) (err error){
	defer conn.Close()

	var bytes [4096]byte
	for {
		// 会一直等待客户端发送数据
		_, err = conn.Read(bytes[:4])
		if err != nil {
			return
		}
		fmt.Println(bytes[:4])
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		go process(conn)
	}

}
