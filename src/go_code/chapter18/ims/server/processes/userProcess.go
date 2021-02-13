package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/common/message"
	"go_code/chapter18/ims/server/model"
	"go_code/chapter18/ims/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
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

	// USE ImsUserDao and validate in redis
	user, err := model.ImsUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginRtnMsg.Code = 500
			loginRtnMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD_INVALID {
			loginRtnMsg.Code = 403
			loginRtnMsg.Error = err.Error()
		} else {
			loginRtnMsg.Code = 505
			loginRtnMsg.Error = "Server Error"
		}
	} else {
		loginRtnMsg.Code = 200
		// update userId in `u` add user into onlineUser list `userMgr`
		(*u).UserId = loginMsg.UserId
		userMgr.AddOnlineUser(u)
		for id, _ := range userMgr.onlineUsers {
			loginRtnMsg.UserIds = append(loginRtnMsg.UserIds, id)
		}
		fmt.Println("Server.UserDaoRtn:", user)
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

func (u *UserProcess) ServerProcessSignup(msg *message.Message) (err error) {
	// 1. get Data, deserialize to LoginMsg
	var signupMsg message.SignupMsg
	err = json.Unmarshal([]byte(msg.Data), &signupMsg)
	if err != nil {
		fmt.Println("Server.process.serverProcessMsg.ServerProcessSignup : Error -", err)
		return
	}

	// declare Message
	var rtnMsg message.Message
	rtnMsg.Type = message.LoginRtnMsgType

	// declare signupRtnMsg and assign value
	var signupRtnMsg message.SignupRtnMsg

	// USE ImsUserDao and validate in redis
	err = model.ImsUserDao.Signup(&signupMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			signupRtnMsg.Code = 500
			signupRtnMsg.Error = err.Error()
		} else {
			signupRtnMsg.Code = 505
			signupRtnMsg.Error = "Server Error"
		}
	} else {
		signupRtnMsg.Code = 200
		fmt.Println("Server : UserDao Signup Success")
	}

	// serialize
	data, err := json.Marshal(signupRtnMsg)
	if err != nil {
		fmt.Println("Server.process.serverProcessMsg.serverProcessSignup : signupRtnMsg Marshall Error -", err)
		return
	}

	// assign data
	rtnMsg.Data = string(data)

	// marshall msg
	data, err = json.Marshal(rtnMsg)
	if err != nil {
		fmt.Println("Server.process.serverProcessMsg.serverProcessSignup : rtnMsg Marshall Error -", err)
		return
	}

	// send data -> writePkg
	transfer := utils.Transfer{
		Conn: (*u).Conn,
	}
	err = transfer.WritePkg(data)

	return
}