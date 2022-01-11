package main

import (
	"chatroom/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	// 界面
	var option int
	for {
		fmt.Println("-------------------待定----------------")
		fmt.Println("1、登录")
		fmt.Println("2、注册")
		fmt.Println("3、退出")
		fmt.Print("输入你的选项：")
		fmt.Scanf("%d", &option)
		switch option {
		case 1 :
			fmt.Println("登录")
			Login()
		case 2 :
			fmt.Println("注册")
		case 3 :
			fmt.Println("退出")
		}
		if option == 3 {
			break
		}
	}
}

// 将id和密码发送给服务器校对
func Login(){
	conn, err := net.Dial("tcp","localhost:1234")
	// 登录信息
	loginMessage := new(message.LoginMessage)
	// 将密码和id设置
	loginMessage.UserId = 123
	loginMessage.Password = "abc"
	// 封装成message
	data := new(message.Message)
	data.Type = message.MESSAGE_TYPE_LOGIN
	messData, err := json.Marshal(loginMessage)
	if err != nil {
		fmt.Printf("json marshal = %v", err)
		return
	}

	// 将长度转化为字节数组
	// 先获得无符号整形长度
	var pkgLen = uint32(len(messData))
	// 定义存储长度的字节数组
	var bytes [4]byte
	// 然后将长度转换为字节数组
	binary.BigEndian.PutUint32(bytes[:4], pkgLen)
	data.Data = string(messData)
	fmt.Println(bytes)
	conn.Write(bytes[:4])
}