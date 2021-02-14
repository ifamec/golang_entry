package processes

import (
	"fmt"
	"go_code/chapter18/ims/common/message"
)

// map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// handle return message
func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg) {

	// if user exist
	user, isExist := onlineUsers[notifyUserStatusMsg.UserId]
	if !isExist {
		user = &message.User{
			UserId:     notifyUserStatusMsg.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status
	onlineUsers[notifyUserStatusMsg.UserId] = user
	printOnlineUser()
}

func printOnlineUser()  {
	fmt.Println("=== Current Online User List:")
	for id, _ := range onlineUsers {
		fmt.Println("=== User Id:", id)
	}
	fmt.Println("==========")
}