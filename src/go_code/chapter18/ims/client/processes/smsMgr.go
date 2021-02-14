package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chapter18/ims/common/message"
)

func printGroupMsg(msg *message.Message) {
	var smsResMsg message.SmsResMsg
	err := json.Unmarshal([]byte(msg.Data), &smsResMsg)
	if err != nil {
		fmt.Println("Server : SmsMgr Unmarshall Error -", err)
		return
	}

	fmt.Println(smsResMsg.User.UserId, ":", smsResMsg.Content)
	fmt.Println()
}
