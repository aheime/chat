package process

import (
	"encoding/json"
	"fmt"
	"gotest/chat/client/utils"
	"gotest/chat/common"
	"net"
)

func ShowMenu() {
	fmt.Println("\t\t\t\t1.显示在线用户列表")
	fmt.Println("\t\t\t\t2.进入聊天室")
	fmt.Println("\t\t\t\t3.信息列表")
	fmt.Println("\t\t\t\t4.退出系统")
	var key int
	var content string
	var loop = true
	smsProcess := &SmsProcess{}
	for loop {
		fmt.Println("请选择[1-4]")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("-----当前在线用户列表------")
			OnlineUserList()
		case 2:
			for true {
				fmt.Scanln(&content)
				smsProcess.sendGroupMsg(content)
			}

		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("退出系统")
			loop = false
			//os.Exit(0)
		default:
			fmt.Println("选择错误")
		}
	}

}

func serverMsg(conn net.Conn) {
	fmt.Println("客户端等待读取服务端消息...")
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
		}

		switch msg.Type {
		case common.NotifyUserStatusMsgType:
			var notifyUserStatusMsg common.NotifyUserStatusMsg
			err := json.Unmarshal([]byte(msg.Data), &notifyUserStatusMsg)
			if err != nil {
				fmt.Println("json err=", err)
			}
			saveUserStatus(&notifyUserStatusMsg, true)

		case common.SmsMsgType:
			smsProcess := &SmsProcess{}
			smsProcess.outShowGroupMsg(&msg)
		default:
			fmt.Println("消息类型错误")
		}
	}
}
