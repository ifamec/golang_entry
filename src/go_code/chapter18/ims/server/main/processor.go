package main

import (
	"fmt"
	"go_code/chapter18/ims/common/message"
	"go_code/chapter18/ims/server/processes"
	"go_code/chapter18/ims/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// handle msg based on msg type
func (p *Processor) serverProcessMsg(msg *message.Message) (err error) {
	// check group msg
	fmt.Println("Msg:", msg)
	switch msg.Type {
	case message.LoginMsgType:
		user := &processes.UserProcess{
			Conn: (*p).Conn,
		}
		err = user.ServerProcessLogin(msg)
	case message.SignupMsgType:
		user := &processes.UserProcess{
			Conn: (*p).Conn,
		}
		err = user.ServerProcessSignup(msg)
	case message.SmsMsgType:
		sms := &processes.SmsProcess{}
		err = sms.ForwardGroupMsg(msg)
	default:
		fmt.Println("Msg Type Not Exist")
	}
	return
}

func (p *Processor) mainProcess() (err error) {

	// read client msg
	for {
		// Transfer object
		transfer := utils.Transfer{
			Conn: (*p).Conn,
		}
		// pack read data as a function readPkg
		msg, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server.process : Client Exit")
				return err
			} else {
				fmt.Println("Server.process : readPkg Error -", err)
				return err
			}
		}
		fmt.Println("Server.process : readPkg Success -", msg)

		err = (*p).serverProcessMsg(&msg)
		if err != nil {
			fmt.Println("Server.process : serverProcessMsg Error -", err)
			return err
		}

	}

}