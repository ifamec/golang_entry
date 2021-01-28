package main

import "fmt"

func main() {
	var selection string
	var exit bool = false
	var balance float64 = 10000
	var delta float64 = 0.0
	var note string = ""
	var detail string = "Action\tBalance\tDelta\tDetails"
	var hasRecord bool = false

	for {
		if exit {
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
		fmt.Scanln(&selection)

		switch selection {
		case "1":
			fmt.Println("-------------- Detail --------------")
			if ! hasRecord {
				fmt.Println("No Record, Add one?")
			} else {
				fmt.Println(detail)
			}
		case "2":
			fmt.Printf("Record Earning: ")
			fmt.Scanln(&delta)
			balance += delta
			fmt.Printf("Earning Details: ")
			fmt.Scanln(&note)
			hasRecord = ! hasRecord
			detail += fmt.Sprintf("\nEarn\t%v\t%v\t%s", delta, balance, note)
		case "3":
			fmt.Printf("Record Expense: ")
			fmt.Scanln(&delta)
			if delta > balance {
				fmt.Printf("Balance not enough")
				break
			}
			balance -= delta
			fmt.Printf("Earning Details: ")
			fmt.Scanln(&note)
			hasRecord = ! hasRecord
			detail += fmt.Sprintf("\nExpense\t%v\t%v\t%s", delta, balance, note)
		case "4":
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
				exit = !exit
			}
		default:  fmt.Println("Error Retry")
		}
	}
}