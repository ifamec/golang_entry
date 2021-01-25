package model

import "fmt"

type Account struct {
	accountNumber string
	balance float64
	password string
}

func NewAccount() *Account {
	return &Account{}
}

func (an *Account) SetAccountNumber(str string)  {
	if len(str) >= 6 && len(str) <= 10 {
		(*an).accountNumber = str
	} else {
		fmt.Println("Account Number Length Should Between 6 - 10")
	}
}
func (an *Account) GetAccountNumber() string {
	return (*an).accountNumber
}

func (an *Account) SetBalance(bal float64)  {
	if bal > 20 {
		(*an).balance = bal
	} else {
		fmt.Println("Balance Should > 20")
	}
}
func (an *Account) GetBalance() float64 {
	return (*an).balance
}

func (an *Account) SetPassword(pwd string)  {
	if len(pwd) == 6 {
		(*an).password = pwd
	} else {
		fmt.Println("Password should be 6 digit")
	}
}
func (an *Account) GetPassword() string {
	return (*an).password
}

func (an *Account) String() string {
	return "ACCOUNT NUMBER: " + (*an).accountNumber + "\tPASSWORD: " + (*an).password + "\tBALANCE: " + fmt.Sprintf("%.2f", (*an).balance)
}