package main

import (
	"fmt"
	"go_code/chapter11/encapsulation01/model"
)

func main()  {
	var p = model.NewPerson("Tom", 30, 4000)
	fmt.Println(*p)

	fmt.Println((*p).Name)
	(*p).SetAge(18)
	fmt.Println((*p).GetAge())
	(*p).SetSalary(60000.00)
	fmt.Println((*p).GetSalary())
}