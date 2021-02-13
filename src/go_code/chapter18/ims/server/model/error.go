package model

import "errors"

// define error per business logic

var (
	ERROR_USER_NOTEXISTS = errors.New("User Not Exist")
	ERROR_USER_EXISTS = errors.New("User ALready Exist")
	ERROR_USER_PWD_INVALID = errors.New("Password Invalid")
)