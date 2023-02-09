package process

import "fmt"

var userMgr *UserMgr

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

func (t *UserMgr) AddOnlineUser(user *UserProcess) {
	t.OnlineUsers[user.UserId] = user
}

func (t *UserMgr) DelOnlineUser(userId int) {
	delete(t.OnlineUsers, userId)
}

func (t *UserMgr) OnlineUserList() map[int]*UserProcess {

	return t.OnlineUsers
}

func (t *UserMgr) GetOnlineUserById(userId int) (user *UserProcess, err error) {

	user, ok := t.OnlineUsers[userId]
	if !ok {
		err = fmt.Errorf("当前%d用户不在线", userId)
	}
	return
}
