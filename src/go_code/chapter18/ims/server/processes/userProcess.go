package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/common/message"
	"go_code/chapter18/ims/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (u *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	// 1. get Data, deserialize to LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("Server.process.serverProcessMsg.serverProcessLogin : Error -", err)
		return
	}

	// declare Message
	var rtnMsg message.Message
	rtnMsg.Type = message.LoginRtnMsgType

	// declare LoginRtnMsg and assign value
	var loginRtnMsg message.LoginRtnMsg
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
		// valid
		loginRtnMsg.Code = 200
	} else {
		// invalid
		loginRtnMsg.Code = 500
		loginRtnMsg.Error = "User InValid, Please SignUp"
	}

	// serialize
	data, err := json.Marshal(loginRtnMsg)
	if err != nil {
		fmt.Println("Server.process.serverProcessMsg.serverProcessLogin : loginRtnMsg Marshall Error -", err)
		return
	}

	// assign data
	rtnMsg.Data = string(data)

	// marshall msg
	data, err = json.Marshal(rtnMsg)
	if err != nil {
		fmt.Println("Server.process.serverProcessMsg.serverProcessLogin : rtnMsg Marshall Error -", err)
		return
	}

	// send data -> writePkg
	transfer := utils.Transfer{
		Conn: (*u).Conn,
	}
	err = transfer.WritePkg(data)

	return
}