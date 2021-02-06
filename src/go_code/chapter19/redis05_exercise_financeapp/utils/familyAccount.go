package utils

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"

	// "github.com/gomodule/redigo/redis"
)

type familyAccount struct {
	selection string
	exit bool
	balance float64
	delta float64
	note string
	detail string
	hasRecord bool
}

// factory mode
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		selection: "",
		exit: false,
		balance: 10000.0,
		delta: 0.0,
		note: "",
		detail: "Action\tDelta\tBalance\tDetails",
		hasRecord: false,
	}
}
func (a *familyAccount) Auth() {
	var trial int
	var id, pwd string
	for {
		if (*a).exit {
			fmt.Println("Exit")
			return
		}
		fmt.Printf("Provide Your ID: ")
		fmt.Scanln(&id)
		fmt.Printf("Provide Your Password: ")
		fmt.Scanln(&pwd)
		if id == "householder" && pwd == "888888" {
			(*a).showMainMenu()
		} else {
			trial++
			if trial >=3 {
				fmt.Println("Please Re-run To Use")
				break
			}
			fmt.Println("Error Retry, You have", 3 - trial, "trial(s)")
		}
	}
}

func (a *familyAccount) showMainMenu() {
	for {
		if (*a).exit {
			fmt.Println("Exit")
			break
		}

		fmt.Println()
		fmt.Println("-------- Family Account App --------")
		fmt.Println("\t 1. Detail")
		fmt.Println("\t 2. Earning")
		fmt.Println("\t 3. Expense")
		fmt.Println("\t 4. Transfer")
		fmt.Println("\t 5. Save & Exit")
		fmt.Println("\t 6. Recover Data")
		fmt.Printf("Please Select < 1 - 6 >: ")
		fmt.Scanln(&(*a).selection)

		switch (*a).selection {
		case "1": (*a).showDetails()
		case "2": (*a).recordEarning()
		case "3": (*a).recordExpense()
		case "4": (*a).transfer()
		case "5": (*a).exitApp()
		case "6": (*a).recoverData()
		default:  fmt.Println("Error Retry")
		}
	}
}

func (a *familyAccount) showDetails() {
	fmt.Println("-------------- Detail --------------")
	if ! (*a).hasRecord {
		fmt.Println("No Record, Add one?")
	} else {
		fmt.Println((*a).detail)
	}
}
func (a *familyAccount) recordEarning() {
	fmt.Printf("Record Earning: ")
	fmt.Scanln(&(*a).delta)
	(*a).balance += (*a).delta
	fmt.Printf("Earning Details: ")
	fmt.Scanln(&(*a).note)
	(*a).hasRecord = true
	(*a).detail += fmt.Sprintf("\nEarn\t%v\t%v\t%s", (*a).delta, (*a).balance, (*a).note)
}
func (a *familyAccount) recordExpense() {
	fmt.Printf("Record Expense: ")
	fmt.Scanln(&(*a).delta)
	if (*a).delta > (*a).balance {
		fmt.Printf("Balance not enough")
		// break
	}
	(*a).balance -= (*a).delta
	fmt.Printf("Earning Details: ")
	fmt.Scanln(&(*a).note)
	(*a).hasRecord = true
	(*a).detail += fmt.Sprintf("\nExpense\t%v\t%v\t%s", (*a).delta, (*a).balance, (*a).note)
}
func (a *familyAccount) transfer() {
	var receiver string
	fmt.Printf("Transfer Amount: ")
	fmt.Scanln(&(*a).delta)
	fmt.Printf("Transfer To: ")
	fmt.Scanln(&receiver)

	if (*a).delta > (*a).balance {
		fmt.Printf("Balance not enough")
		// break
	}
	(*a).balance -= (*a).delta
	(*a).hasRecord = true
	(*a).detail += fmt.Sprintf("\nTransfer\t%v\t%v\t%s", (*a).delta, (*a).balance, receiver)
}
func (a *familyAccount) exitApp() {
	fmt.Printf("Are You Sure To Save & Exit (Y/N): ")
	var confirm string
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" || confirm == "n" || confirm == "N" {
			break
		}
		fmt.Printf("Input Error, Are You Sure To Exit (Y/N): ")
	}
	if confirm == "y" || confirm == "Y"{
		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			fmt.Println("Redis Conn Error", err)
		}
		defer conn.Close()

		_, err = conn.Do("HMSET", "familyAccount", "Delta", (*a).delta, "Balance", (*a).balance, "Details", (*a).detail)
		if err != nil {
			fmt.Println("Redis HMSET Error", err)
		}
		fmt.Println("Save Success, Exit...")
		(*a).exit = true
	}
}
func (a *familyAccount) recoverData() {
	fmt.Printf("Are You Sure To Recover Data (Y/N): ")
	var confirm string
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" || confirm == "n" || confirm == "N" {
			break
		}
		fmt.Printf("Input Error, Are You Sure Recover Data (Y/N): ")
	}
	if confirm == "y" || confirm == "Y"{
		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			fmt.Println("Redis Conn Error", err)
		}
		defer conn.Close()

		values, err := redis.Strings(conn.Do("HMGET", "familyAccount", "Delta", "Balance", "Details"))
		if err != nil {
			fmt.Println("Redis HMSET Error", err)
		}
		(*a).delta, _ = strconv.ParseFloat(values[0], 64)
		(*a).balance, _ = strconv.ParseFloat(values[1], 64)
		(*a).detail = values[2]
		(*a).hasRecord = true
	}
}