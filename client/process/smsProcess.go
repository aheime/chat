package process

import (
	"encoding/json"
	"fmt"
	"gotest/chat/client/utils"
	"gotest/chat/common"
)

type SmsProcess struct {
}

func (t *SmsProcess) sendGroupMsg(content string) (err error) {

	var smsMsg common.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = CurUser.UserId
	smsMsg.UserName = CurUser.UserName

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}

	var msg common.Message
	msg.Type = common.SmsMsgType
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)

	return
}

func (t *SmsProcess) outShowGroupMsg(msg *common.Message) {

	var smsMsg common.SmsMsg

	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}
	if CurUser.UserId == smsMsg.UserId {
		fmt.Printf("\t\t\t\t我：%v\n", smsMsg.Content)
	} else {
		fmt.Printf("%v：%v\n", smsMsg.UserName, smsMsg.Content)
	}
}
