package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/client/utils"
	"go_code/chapter18/ims/common/message"
	"net"
	"os"
)

func ShowMenu() {
	var key int
	var content string
	fmt.Println("-------- Hello XXX --------")
	fmt.Println("\t 1. Online User List")
	fmt.Println("\t 2. Send Message")
	fmt.Println("\t 3. History")
	fmt.Println("\t 4. Exit")
	fmt.Print("--------  Select: ")
	fmt.Scanln(&key)

	// smsMsg
	smsProcess := &SmsProcess{}

	switch key {
	case 1:
		printOnlineUser()
	case 2:
		fmt.Print("Message: ")
		fmt.Scanln(&content)
		err := smsProcess.SendGroupMsg(content)
		if err != nil {
			fmt.Println("Client : Send Group Msg Error", err)
		}
	case 3:
		fmt.Println("History")
	case 4:
		fmt.Println("Exit System")
		os.Exit(0)
	default:
		fmt.Println("Input Invalid, Retry")
	}
}

func msgPush(conn net.Conn) {
	transfer := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("Client : Wait For Server Response")
		msg, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("Client : Server Read Error -", err)
			return
		}
		// show message
		fmt.Println("Client : Msg From Server -", msg)
		switch msg.Type {
		case message.NotifyUserStatusMsgType: // client online
			// 1. get notify message
			var notifyUserStatusMsg message.NotifyUserStatusMsg
			err = json.Unmarshal([]byte(msg.Data), &notifyUserStatusMsg)
			if err != nil {
				fmt.Println("Client : NotifyUserStatusMsgType Unmarshall Error -", err)
				return
			}
			// 2. add to client side online user map[int]User
			updateUserStatus(&notifyUserStatusMsg)
		case message.SmsResMsgType: // client online
			printGroupMsg(&msg)
		default:
			fmt.Println("Unknown Msg Type")
		}
	}
}
