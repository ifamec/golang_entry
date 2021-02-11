package processes

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/client/utils"
	"go_code/chapter18/ims/common/message"
	"net"
)

type UserProcess struct {
}

func (u *UserProcess) Login(userId int, userPwd string) (err error) {
	// fmt.Println("Your Input:", userId, userPwd)
	// return nil

	// 1 connect to server
	conn, err := net.Dial("tcp", "localhost:8889") // read config in future
	if err != nil {
		fmt.Println("Client : Dail Error -", err)
		return err
	}
	defer conn.Close()

	// 2 conn send msg to server
	var msg message.Message
	msg.Type = message.LoginMsgType

	// 3 create login message
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	// 4 serialize loginMsg
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("Client : Login Msg json.Marshall Error -", err)
		return err
	}

	// 5 assign to msg.Data
	msg.Data = string(data)

	// 6 serialize msg
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("Client : Msg json.Marshall Error -", err)
		return err
	}

	// 7 send data
	// 7.1 send data length
	// get length then to []byte
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("Client : Conn Error -", err)
	}
	fmt.Println("Client : Msg Len Sent Success,", n, "Bytes Sent")

	// 7.2 send data
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Client : Conn Error -", err)
	}
	// fmt.Println("Client : Msg Sent Success, Len", len(data), "Data:", string(data))
	fmt.Println("Client : Msg Sent Success")

	// 8 handle server response
	transfer := &utils.Transfer{
		Conn: conn,
	}

	msg, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("Client : readPkg Error -", err)
		return
	}

	var loginRtnMsg message.LoginRtnMsg
	err = json.Unmarshal([]byte(msg.Data), &loginRtnMsg)
	if err != nil {
		fmt.Println("Client : Unmarshall Server Response Error -", err)
		return
	}
	if loginRtnMsg.Code == 200 {
		// fmt.Println("Login Success")
		// another goroutine for server connection if anything pushed show
		go msgPush(conn)
		// show main menu
		for {
			ShowMenu()
		}
	} else if loginRtnMsg.Code == 500 {
		fmt.Println(loginRtnMsg.Error)
	}

	return
}
