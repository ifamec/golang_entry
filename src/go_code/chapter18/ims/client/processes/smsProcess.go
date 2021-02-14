package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/client/utils"
	"go_code/chapter18/ims/common/message"
)

type SmsProcess struct {
}

func (sms *SmsProcess) SendGroupMsg(content string) (err error) {

	// 1. msg
	var msg message.Message
	msg.Type = message.SmsMsgType

	// 2. smsMsg
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.User.UserId = CurUser.UserId
	smsMsg.User.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("Client : SMS Msg json.Marshall Error -", err)
		return
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("Client : Msg json.Marshall Error -", err)
		return
	}

	transfer := utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("Client : SMSMsg SendGroup json.Marshall Error -", err)
		return
	}

	return
}