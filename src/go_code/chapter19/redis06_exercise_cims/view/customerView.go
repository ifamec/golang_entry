package main

import (
	"fmt"
	"go_code/chapter19/redis06_exercise_cims/model"
	"go_code/chapter19/redis06_exercise_cims/service"
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
		fmt.Println("\t\t5. Query Customer")
		fmt.Println("\t\t6. Exit System")
		fmt.Println("\t\t7. Save Data")
		fmt.Println("\t\t8. Recover Data")
		fmt.Print("Select < 1 - 8 >: ")
		fmt.Scanln(&(*c).selection)

		switch (*c).selection {
		case "1": (*c).addCustomer()
		case "2": (*c).modifyCustomer()
		case "3": (*c).deleteCustomer()
		case "4": (*c).showList()
		case "5": (*c).queryCustomer()
		case "6": (*c).exitApp()
		case "7": (*c).saveData()
		case "8": (*c).recoverData()
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
	fmt.Printf("Please Provide Id (-1 to exit): ")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("Cancel Delete")
		return
	}
	if (*c).customerService.FindById(id) == -1 {
		fmt.Println("ID Not Exist")
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
func (c *customerView) modifyCustomer() {
	var isSuccess bool
	var id, age int = -1, 0
	var name, gender, phone, email string
	var existingCustomer model.Customer
	fmt.Println("-------------------------- Modify Customer -------------------------")
	fmt.Printf("Please Provide Id (-1 to exit): ")
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("Cancel Modify")
		return
	}
	if (*c).customerService.FindById(id) == -1 {
		fmt.Println("ID Not Exist")
		return
	}
	existingCustomer = (*c).customerService.List()[(*c).customerService.FindById(id)]
	name, gender, age, phone, email = existingCustomer.Name, existingCustomer.Gender, existingCustomer.Age, existingCustomer.Phone, existingCustomer.Email
	fmt.Println("Hit Enter To Preserve Current Value")
	fmt.Printf("Id: (%v)\n", id)
	fmt.Printf("Name(%v): ", name)
	fmt.Scanln(&name)
	fmt.Printf("Gender(%v): ", gender)
	fmt.Scanln(&gender)
	fmt.Printf("Age(%v): ", age)
	fmt.Scanln(&age)
	fmt.Printf("Phone(%v): ", phone)
	fmt.Scanln(&phone)
	fmt.Printf("Email(%v): ", email)
	fmt.Scanln(&email)
	fmt.Println("--------------------- Modified Data Collected ----------------------")
	var modifiedCustomer = model.NewCustomer(id, name, gender, age, phone, email)
	fmt.Printf("Are You Sure To Modify (Y/N): ")
	for {
		var confirm string = ""
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" {
			isSuccess = (*c).customerService.Modify(modifiedCustomer)
			if isSuccess {
				fmt.Println("--------------------- Modify Customer Success ----------------------")
			} else {
				fmt.Println("---------------------- Modify Customer Error -----------------------")
			}
			break
		} else if confirm == "n" || confirm == "N" {
			fmt.Println("Cancel Modify")
			break
		}
		fmt.Printf("Input Error, Are You Sure To Delete (Y/N): ")
	}
}
func (c *customerView) queryCustomer() {
	fmt.Println("Not Available.")
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
		(*c).exit = true
	}
}
func (c *customerView) saveData() {
	fmt.Printf("Are You Sure To Recover Data (Y/N): ")
	var confirm string
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" || confirm == "n" || confirm == "N" {
			break
		}
		fmt.Printf("Input Error, Are You Sure To Recover Data (Y/N): ")
	}
	if confirm == "y" || confirm == "Y"{
		(*c).customerService.SaveData()
	}
}
func (c *customerView) recoverData() {
	fmt.Printf("Are You Sure To Recover Data (Y/N): ")
	var confirm string
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" || confirm == "n" || confirm == "N" {
			break
		}
		fmt.Printf("Input Error, Are You Sure To Recover Data (Y/N): ")
	}
	if confirm == "y" || confirm == "Y"{
		(*c).customerService.RecoverData()
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
