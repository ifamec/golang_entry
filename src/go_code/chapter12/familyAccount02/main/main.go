package main

import (
	"fmt"
	"go_code/chapter12/familyAccount02/utils"
)

func main()  {
	fmt.Println("OOP")
	utils.NewFamilyAccount().ShowMainMenu()
}