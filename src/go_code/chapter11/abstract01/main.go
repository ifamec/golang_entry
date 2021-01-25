package main

import "fmt"

type Account struct {
	AccountNumber string
	Password      string
	Balance       float64
}

func (a *Account) Deposit(amount float64, pwd string) {
	// validate pwd
	if (*a).Password != pwd {
		fmt.Println("PWD ERROR")
		return
	}

	// check deposit amount
	if amount < 0 {
		fmt.Println("AMOUNT ERROR")
		return
	}

	(*a).Balance += amount
	fmt.Println("Deposit SUCCESS")
}

func (a *Account) Withdraw(amount float64, pwd string) {
	// validate pwd
	if (*a).Password != pwd {
		fmt.Println("PWD ERROR")
		return
	}

	// check deposit amount
	if amount <= 0 || amount > (*a).Balance {
		fmt.Println("AMOUNT ERROR")
		return
	}

	(*a).Balance -= amount
	fmt.Println("Withdraw SUCCESS")
}

func (a *Account) CheckBalance(pwd string) {
	// validate pwd
	if (*a).Password != pwd {
		fmt.Println("PWD ERROR")
		return
	}
	fmt.Println("YOUR ACCOUNT",(*a).AccountNumber,"BALANCE IS:", (*a).Balance)
}

func main() {
	var user Account = Account{"C00001", "888888", 10000.64}
	user.CheckBalance("888888")
	user.Deposit(5000.82, "888888")
	user.CheckBalance("888888")
	user.Withdraw(1325.43, "888888")
	user.CheckBalance("888888")
}
