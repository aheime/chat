package common

const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMsg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResMsgType      = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType              = "SmsMsg"
)

const (
	UserOffline = iota
	UserOnline
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResMsg struct {
	Code        int              `json:"code"`
	OnlineUsers []map[int]string `json:"online_users"`
	UserInfo    map[string]string
	Error       string `json:"error"`
}

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMsg struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Status   int    `json:"status"`
}

type SmsMsg struct {
	Content string `json:"content"`
	User
}
