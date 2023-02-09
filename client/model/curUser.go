package model

import (
	"gotest/chat/common"
	"net"
)

type CurUser struct {
	Conn net.Conn
	common.User
}
