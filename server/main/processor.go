package main

import (
	"fmt"
	"gotest/chat/common"
	"gotest/chat/server/process"
	"gotest/chat/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (t *Processor) serverProcessMsg(msg *common.Message) (err error) {

	switch msg.Type {
	case common.LoginMsgType:
		userProcess := &process.UserProcess{
			Conn: t.Conn,
		}
		err = userProcess.ServerProcessLogin(msg)
	case common.RegisterMsgType:
		userProcess := &process.UserProcess{
			Conn: t.Conn,
		}
		err = userProcess.ServerProcessRegister(msg)
	case common.SmsMsgType:
		smsProcess := &process.SmsProcess{}
		smsProcess.SendGroupMsg(msg)
	default:
		fmt.Println("消息类型错误")
	}
	return
}

func (t *Processor) run() (err error) {
	for {

		tf := &utils.Transfer{
			Conn: t.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出")
			} else {
				fmt.Println("read pkg err=", err)
			}
			return err
		}

		err = t.serverProcessMsg(&msg)
		if err != nil {
			return err
		}

	}
}
