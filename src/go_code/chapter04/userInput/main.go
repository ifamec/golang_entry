package main

import "fmt"

func main()  {
	var name string
	var age byte
	var salary float32
	var isPass bool

	// Scanln
	fmt.Printf("Input Name:")
	fmt.Scanln(&name) // Wait for user input and hit enter
	fmt.Printf("Input Age:")
	fmt.Scanln(&age)
	fmt.Printf("Input Salary:")
	fmt.Scanln(&salary)
	fmt.Printf("Input isPass:")
	fmt.Scanln(&isPass)

	fmt.Printf("name:%v\tage:%v\tsalary:%v\tisPass:%v\n", name, age, salary, isPass)
	// Scanf
	fmt.Printf("Input \"Name Age Salary isPass\" separate with space\n")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)

	fmt.Printf("name:%v\tage:%v\tsalary:%v\tisPass:%v\n", name, age, salary, isPass)
}