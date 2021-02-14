package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/common/message"
	"go_code/chapter18/ims/server/utils"
	"net"
)

type SmsProcess struct {

}

func (sp *SmsProcess) ForwardGroupMsg(msg *message.Message) (err error) {

	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("Server : SmsProcess Unmarshall Error -", err)
		return
	}

	var msgRes message.Message
	msgRes.Type = message.SmsResMsgType

	var smsResMsg message.SmsResMsg
	smsResMsg.Content = smsMsg.Content
	smsResMsg.User = smsMsg.User

	data, err := json.Marshal(smsResMsg)
	if err != nil {
		fmt.Println("Server : SmsProcess smsResMsg Marshall Error -", err)
		return
	}
	msgRes.Data = string(data)

	data, err = json.Marshal(msgRes)
	if err != nil {
		fmt.Println("Server : SmsProcess Msg Marshall Error -", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		// exclude the sender
		if id == smsMsg.User.UserId {
			continue
		}
		(*sp).SendMsgToAllOnlineUser(data, up.Conn)
	}

	return
}

func (sp *SmsProcess) SendMsgToAllOnlineUser(data []byte, conn net.Conn)  {
	transfer := utils.Transfer{
		Conn: conn,
	}
	err := transfer.WritePkg(data)
	if err != nil {
		fmt.Println("Server : SmsProcess Forward Error -", err)
		return
	}
}