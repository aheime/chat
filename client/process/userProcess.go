package process

import (
	"encoding/json"
	"fmt"
	"gotest/chat/client/utils"
	"gotest/chat/common"
	"net"
	"strconv"
)

type UserProcess struct {
}

func (t *UserProcess) Register(userName string, password string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8899")

	if err != nil {
		fmt.Println("客户端连接错误", err)
		return
	}

	defer conn.Close()

	var registerMsg common.RegisterMsg
	registerMsg.User.UserName = userName
	registerMsg.User.Password = password

	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	msgStruct := common.Message{
		Type: common.RegisterMsgType,
		Data: string(data),
	}

	data, err = json.Marshal(msgStruct)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注册数据失败err=", err)
		return
	}

	msg, err := common.ReadPkg(conn)

	var registerResMsg common.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)

	if registerResMsg.Code == 200 {
		fmt.Println("注册成功")
	} else {
		fmt.Println("注册失败err=", registerResMsg.Error)
	}

	return
}

func (t *UserProcess) Login(userName string, password string) (err error) {

	conn, err := net.Dial("tcp", "127.0.0.1:8899")

	if err != nil {
		fmt.Println("客户端连接错误", err)
		return
	}

	defer conn.Close()

	loginMsg := common.LoginMsg{
		UserName: userName,
		Password: password,
	}

	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	msgStruct := common.Message{
		Type: common.LoginMsgType,
		Data: string(data),
	}

	data, err = json.Marshal(msgStruct)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	tf := &utils.Transfer{Conn: conn}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送登录数据失败err=", err)
		return
	}

	msg, err := common.ReadPkg(conn)

	var loginResMsg common.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)

	if loginResMsg.Code == 200 {

		CurUser.Conn = conn
		CurUser.UserId, _ = strconv.Atoi(loginResMsg.UserInfo["UserId"])
		CurUser.UserName = loginResMsg.UserInfo["UserName"]
		for _, v := range loginResMsg.OnlineUsers {
			for id, name := range v {
				if strconv.Itoa(id) == loginResMsg.UserInfo["UserId"] {
					continue
				}
				notifyUserStatusMsg := &common.NotifyUserStatusMsg{
					UserId:   id,
					UserName: name,
					Status:   common.UserOnline,
				}
				saveUserStatus(notifyUserStatusMsg, false)
			}

		}
		fmt.Printf("\t\t\t\t--欢迎%v登录成功--\n", loginResMsg.UserInfo["UserName"])

		go serverMsg(conn)

		ShowMenu()
	} else {
		fmt.Println("登录失败err=", loginResMsg.Error)
	}

	return
}
