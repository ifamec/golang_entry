package model

import (
	"go_code/chapter18/ims/common/message"
	"net"
)

type CurUser struct { // global
	Conn net.Conn
	message.User
}