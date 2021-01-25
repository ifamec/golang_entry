package main

import (
	"fmt"
	"go_code/chapter11/encapsulation02_exercise/model"
)

func main()  {
	var user = model.NewAccount()

	(*user).SetAccountNumber("c00001")
	(*user).SetBalance(123123.12)
	(*user).SetPassword("123123")

	fmt.Println(user)

	fmt.Println((*user).GetAccountNumber())
	fmt.Println((*user).GetBalance())
	fmt.Println((*user).GetPassword())
}