package main

import (
	"fmt"
	"gotest/chat/client/process"
)

func main() {
	var key int

	var loop = true

	for loop {

		fmt.Println("\t\t\t\t 欢迎登录聊天系统")
		fmt.Println("\t\t\t\t 1 登录账号")
		fmt.Println("\t\t\t\t 2 注册账号")
		fmt.Println("\t\t\t\t 3 退出系统")
		fmt.Println("\t\t\t\t 请选择")

		fmt.Scanf("%d\n", &key)

		var useName string
		var pwd string
		switch key {
		case 1:
			fmt.Println("请输入名称")
			fmt.Scanln(&useName)
			fmt.Println("请输入密码")
			fmt.Scanln(&pwd)
			userProcess := &process.UserProcess{}
			userProcess.Login(useName, pwd)
		case 2:
			fmt.Println("请输入名称")
			fmt.Scanln(&useName)
			fmt.Println("请输入密码")
			fmt.Scanln(&pwd)
			userProcess := &process.UserProcess{}
			userProcess.Register(useName, pwd)
		case 3:
		default:
			fmt.Println("你的输入有误，请重新输入")

		}
	}
}
