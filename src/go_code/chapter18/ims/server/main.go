package main

import (
	"encoding/binary"
	"encoding/json"
	_ "errors"
	"fmt"
	"go_code/chapter18/ims/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 4096)
	fmt.Println("Server.process.readPkg : Wait For Client Message")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("Server.process.readPkg : Conn Read Msg Header Error -", err)
		// err = errors.New("read msg header error")
		return
	}
	// fmt.Println("Client.process : Read Success -", buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// read content
	n, err := conn.Read(buf[:pkgLen]) // conn.Read could be blocked only if both side is not closed
	if n != int(pkgLen) || err != nil {
		fmt.Println("Server.process.readPkg : Conn Read Msg Body Error -", err)
		// err = errors.New("read msg body error")
		return
	}
	// pkgLen to message.Message
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("Server.process.readPkg : Unmarshall Msg Error -", err)
		return
	}

	return
}
func writePkg(conn net.Conn, data []byte) (err error) {
	// 1. send length
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("Server : Conn Error -", err)
	}
	fmt.Println("Server : Msg Len Sent Success,", n, "Bytes Sent")

	// 7.2 send data
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("Server : Conn Error -", err)
	}
	// fmt.Println("Client : Msg Sent Success, Len", len(data), "Data:", string(data))
	fmt.Println("Server : Msg Sent Success")
	return
}
// server process login
func serverProcessLogin(conn net.Conn, msg *message.Message) (err error) {
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
	err = writePkg(conn, data)

	return
}

// handle msg based on msg type
func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		err = serverProcessLogin(conn, msg)
	case message.SignupMsgType:
	default:
		fmt.Println("Msg Type Not Exist")
	}
	return
}
// communication
func process(conn net.Conn) {
	// read client message
	defer conn.Close()

	// read client msg
	for {
		// pack read data as a function readPkg
		msg, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server.process : Client Exit")
				return
			} else {
				fmt.Println("Server.process : readPkg Error -", err)
				return
			}
		}
		fmt.Println("Server.process : readPkg Success -", msg)

		err = serverProcessMsg(conn, &msg)
		if err != nil {
			fmt.Println("Server.process : serverProcessMsg Error -", err)
			return
		}
	}
}

func main() {
	// listen port at 8889

	fmt.Println("Server : Start Listen At 8889")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("Server : Listen Error -", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("Server : Wait For Client")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Server : Listen Accept Error -", err)
		}

		// start goroutine if conn success
		go process(conn)
	}
}
