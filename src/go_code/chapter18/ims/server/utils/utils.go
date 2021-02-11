package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/common/message"
	"net"
)

// struct
type Transfer struct {
	Conn net.Conn   // connection
	Buf [4096]byte  // buffer in transfer
}

func (t *Transfer) ReadPkg() (msg message.Message, err error) {
	// buf := make([]byte, 4096)
	fmt.Println("Server.process.readPkg : Wait For Client Message")
	_, err = (*t).Conn.Read((*t).Buf[:4])
	if err != nil {
		fmt.Println("Server.process.readPkg : Conn Read Msg Header Error -", err)
		// err = errors.New("read msg header error")
		return
	}
	// fmt.Println("Client.process : Read Success -", buf[:4])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32((*t).Buf[:4])

	// read content
	n, err := (*t).Conn.Read((*t).Buf[:pkgLen]) // conn.Read could be blocked only if both side is not closed
	if n != int(pkgLen) || err != nil {
		fmt.Println("Server.process.readPkg : Conn Read Msg Body Error -", err)
		// err = errors.New("read msg body error")
		return
	}
	// pkgLen to message.Message
	err = json.Unmarshal((*t).Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("Server.process.readPkg : Unmarshall Msg Error -", err)
		return
	}

	return
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	// 1. send length
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32((*t).Buf[:4], pkgLen)
	n, err := (*t).Conn.Write((*t).Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("Server : Conn Error -", err)
	}
	fmt.Println("Server : Msg Len Sent Success,", n, "Bytes Sent")

	// 7.2 send data
	n, err = (*t).Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("Server : Conn Error -", err)
	}
	// fmt.Println("Client : Msg Sent Success, Len", len(data), "Data:", string(data))
	fmt.Println("Server : Msg Sent Success")
	return
}