package process

import (
	"encoding/json"
	"fmt"
	"gotest/chat/client/model"
	"gotest/chat/common"
	"gotest/chat/server/utils"
	"net"
	"strconv"
)

type UserProcess struct {
	Conn     net.Conn
	UserId   int
	UserName string
}

func (t *UserProcess) ServerProcessLogin(msg *common.Message) (err error) {

	var loginMsg common.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}

	var loginResMsg common.LoginResMsg
	user, err := model.MyUserDao.Login(loginMsg.UserName, loginMsg.Password)
	if err != nil {
		loginResMsg.Code = 500
		loginResMsg.Error = err.Error()

	} else {
		t.UserId = user.UserId
		t.UserName = user.UserName
		userMgr.AddOnlineUser(t)

		loginResMsg.Code = 200
		userInfo := make(map[string]string)
		userInfo["UserId"] = strconv.Itoa(user.UserId)
		userInfo["UserName"] = user.UserName
		loginResMsg.UserInfo = userInfo
		//loginResMsg.OnlineUsers = make([]map[int]string, 10)
		for id, val := range userMgr.OnlineUsers {
			idName := make(map[int]string)
			idName[id] = val.UserName
			loginResMsg.OnlineUsers = append(loginResMsg.OnlineUsers, idName)
		}
	}

	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return err
	}
	msg.Type = common.LoginResMsgType
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return err
	}
	tf := &utils.Transfer{
		Conn: t.Conn,
	}
	err = tf.WritePkg(data)

	err = t.NotifyOtherOnlineUser()
	if err != nil {
		return err
	}

	return err
}

func (t *UserProcess) ServerProcessRegister(msg *common.Message) (err error) {

	var registerMsg common.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return
	}

	var registerResMsg common.RegisterResMsg
	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		registerResMsg.Code = 500
		registerResMsg.Error = err.Error()

	} else {
		registerResMsg.Code = 200
	}

	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return err
	}
	msg.Type = common.RegisterResMsgType
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return err
	}
	tf := &utils.Transfer{
		Conn: t.Conn,
	}
	err = tf.WritePkg(data)

	return err
}

// NotifyOtherOnlineUser 通知其他在线用户
func (t *UserProcess) NotifyOtherOnlineUser() (err error) {

	for id, user := range userMgr.OnlineUsers {
		if id == t.UserId {
			continue
		}
		t.Notify(user)
	}
	return
}

func (t *UserProcess) Notify(user *UserProcess) (err error) {

	notifyStatusMsg := common.NotifyUserStatusMsg{
		UserId:   t.UserId,
		UserName: t.UserName,
		Status:   common.UserOnline,
	}

	data, err := json.Marshal(notifyStatusMsg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return err
	}
	msg := common.Message{
		Type: common.NotifyUserStatusMsgType,
		Data: string(data),
	}
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化错误", err)
		return err
	}
	tf := &utils.Transfer{Conn: user.Conn}
	err = tf.WritePkg(data)

	return
}
