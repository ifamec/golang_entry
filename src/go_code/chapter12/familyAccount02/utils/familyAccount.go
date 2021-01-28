package utils

import "fmt"

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
		detail: "Action\tBalance\tDelta\tDetails",
		hasRecord: false,
	}
}

func (a *familyAccount) ShowMainMenu() {
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
		fmt.Println("\t 4. Exit")
		fmt.Printf("Please Select < 1 - 4 >: ")
		fmt.Scanln(&(*a).selection)

		switch (*a).selection {
		case "1": (*a).showDetails()
		case "2": (*a).recordEarning()
		case "3": (*a).recordExpense()
		case "4": (*a).exitApp()
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
	(*a).hasRecord = ! (*a).hasRecord
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
	(*a).hasRecord = ! (*a).hasRecord
	(*a).detail += fmt.Sprintf("\nExpense\t%v\t%v\t%s", (*a).delta, (*a).balance, (*a).note)
}
func (a *familyAccount) exitApp() {
	fmt.Printf("Are You Sure To Exit (Y/N): ")
	var confirm string
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" || confirm == "n" || confirm == "N" {
			break
		}
		fmt.Printf("Input Error, Are You Sure To Exit (Y/N): ")
	}
	if confirm == "y" || confirm == "Y"{
		(*a).exit = !(*a).exit
	}
}