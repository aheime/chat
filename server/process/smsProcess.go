package process

import (
	"encoding/json"
	"fmt"
	"gotest/chat/common"
	"gotest/chat/server/utils"
	"net"
)

type SmsProcess struct {
}

func (t *SmsProcess) SendGroupMsg(msg *common.Message) (err error) {

	var smsMsg common.SmsMsg

	err = json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json err=", err)
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json err=", err)
	}

	for _, user := range userMgr.OnlineUsers {
		//不发自己
		//if id == smsMsg.UserId {
		//	continue
		//}
		err = t.sendMsgToEachUser(data, user.Conn)
	}
	return
}

func (t *SmsProcess) sendMsgToEachUser(data []byte, conn net.Conn) (err error) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)

	return
}
