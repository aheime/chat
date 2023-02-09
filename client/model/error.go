package model

import "errors"

var (
	ERROR_USER_NOT_EXITSTS = errors.New("用戶不存在")
	ERROR_USER_EXITSTS     = errors.New("用戶存在")
	ERROR_USER_PWD         = errors.New("用戶密码错误")
)
