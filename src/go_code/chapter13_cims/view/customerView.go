package main

import "fmt"

type customerView struct {
	selection string
	exit bool
}
func (c *customerView) mainMenu() {
	for {
		fmt.Println()
		fmt.Println("-------------- CIMS --------------")
		fmt.Println("\t1. Add Customer")
		fmt.Println("\t2. Modify Customer")
		fmt.Println("\t3. Delete Customer")
		fmt.Println("\t4. Customer List")
		fmt.Println("\t5. Exit")
		fmt.Print("Select < 1 - 5 >: ")
		fmt.Scanln(&(*c).selection)

		switch (*c).selection {
		case "1": fmt.Println("ADD")
		case "2": fmt.Println("Modify")
		case "3": fmt.Println("Delete")
		case "4": fmt.Println("List")
		case "5": (*c).exit = true
		default: fmt.Println("Error, Retry")
		}

		if (*c).exit {
			break
		}
	}
	fmt.Println("Exit Successfully, Thank You For Using.")
}

func main() {
	var customerView = customerView{
		selection: "",
		exit: false,
	}

	customerView.mainMenu()
}





