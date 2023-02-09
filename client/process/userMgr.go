package process

import (
	"fmt"
	"gotest/chat/client/model"
	"gotest/chat/common"
)

var onlineUsers = make(map[int]*common.User, 10)
var CurUser model.CurUser

func OnlineUserList() {
	for _, user := range onlineUsers {
		fmt.Printf("用户：%v\n", user.UserName)
	}
}
func saveUserStatus(notifyUserStatusMsg *common.NotifyUserStatusMsg, notify bool) {

	user, ok := onlineUsers[notifyUserStatusMsg.UserId]

	if !ok {
		user = &common.User{
			UserId:   notifyUserStatusMsg.UserId,
			UserName: notifyUserStatusMsg.UserName,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status
	onlineUsers[notifyUserStatusMsg.UserId] = user
	if notify {
		fmt.Printf("用户:%v上线\n", notifyUserStatusMsg.UserName)
	}
}
