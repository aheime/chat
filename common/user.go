package common

type User struct {
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	UserStatus int    `json:"user_status"`
}
