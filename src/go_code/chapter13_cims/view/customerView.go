package main

import (
	"fmt"
	"go_code/chapter13_cims/model"
	"go_code/chapter13_cims/service"
)

type customerView struct {
	selection string
	exit bool
	customerService *service.CustomerService
}

func (c *customerView) mainMenu() {
	for {
		fmt.Println()
		fmt.Println("-------------- Customer Information Management System --------------")
		fmt.Println("\t\t1. Add Customer")
		fmt.Println("\t\t2. Modify Customer")
		fmt.Println("\t\t3. Delete Customer")
		fmt.Println("\t\t4. Customer List")
		fmt.Println("\t\t5. Exit System")
		fmt.Print("Select < 1 - 5 >: ")
		fmt.Scanln(&(*c).selection)

		switch (*c).selection {
		case "1": (*c).addCustomer()
		case "2": fmt.Println("Modify")
		case "3": (*c).deleteCustomer()
		case "4": (*c).showList()
		case "5": (*c).exitApp()
		default: fmt.Println("Error, Retry")
		}

		if (*c).exit {
			break
		}
	}
	fmt.Println("Exit Successfully, Thank You For Using.")
}
func (c *customerView) showList() {
	var customers = (*c).customerService.List()
	fmt.Println("-------------------------- Customers List --------------------------")
	fmt.Println("ID\tName\tGender\tAge\tPhone\t\tEmail")
	for _, v := range customers {
		fmt.Println(v.GetInfo())
	}
	fmt.Println("------------------------ Customers List End ------------------------")
}
func (c *customerView) addCustomer() {
	var isSuccess bool
	var name, gender, phone, email string
	var age int
	fmt.Println("--------------------------- Add Customer ---------------------------")
	fmt.Printf("Name: ")
	fmt.Scanln(&name)
	fmt.Printf("Gender: ")
	fmt.Scanln(&gender)
	fmt.Printf("Age: ")
	fmt.Scanln(&age)
	fmt.Printf("Phone: ")
	fmt.Scanln(&phone)
	fmt.Printf("Email: ")
	fmt.Scanln(&email)
	var customer = model.NewCustomerExceptId(name, gender, age, phone, email)
	isSuccess = (*c).customerService.Add(customer)
	if isSuccess {
		fmt.Println("----------------------- Add Customer Success -----------------------")
	}
}
func (c *customerView) deleteCustomer() {
	var isSuccess bool
	var id int = -1
	fmt.Println("-------------------------- Delete Customer -------------------------")
	fmt.Printf("Please Provide Id: ")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("Cancel Delete")
		return
	}
	fmt.Printf("Are You Sure To Delete? (Y/N): ")
	var confirm string
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" {
			isSuccess = (*c).customerService.Delete(id)
			if isSuccess {
				fmt.Println("--------------------- Delete Customer Success ----------------------")
			} else {
				fmt.Println("ID Not Exist")
				fmt.Println("---------------------- Delete Customer Error -----------------------")
			}
			break
		} else if confirm == "n" || confirm == "N" {
			fmt.Println("Cancel Delete")
			break
		}
		fmt.Printf("Input Error, Are You Sure To Delete (Y/N): ")
	}
}
func (c *customerView) exitApp() {
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
		(*c).exit = !(*c).exit
	}
}

func main() {
	var ui = customerView{
		selection: "",
		exit: false,
		customerService: service.NewCustomerService(),
	}

	ui.mainMenu()
}
