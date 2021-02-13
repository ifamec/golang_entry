package message

// message type const

const (
	LoginMsgType     = "LoginMsg"
	LoginRtnMsgType  = "LoginRtnMsg"
	SignupMsgType    = "SignupMsg"
	SignupRtnMsgType = "SignupRtnMsg"
)

type Message struct {
	Type string `json:"type"` // message type
	Data string `json:"data"` // data
}

// define two message
type LoginMsg struct {
	UserId   int    `json:"userId"`   // user id
	UserPwd  string `json:"userPwd"`  // user password
	UserName string `json:"userName"` // user name
}

type LoginRtnMsg struct {
	Code  int    `json:"code"`  // 500 not registered, 200 success
	Error string `json:"error"` // Error message
}

type SignupMsg struct {
	User User `json:"user"`
}

type SignupRtnMsg struct {
	Code  int    `json:"code"`  // 200 success, 400 user exist
	Error string `json:"error"` // Error message
}
