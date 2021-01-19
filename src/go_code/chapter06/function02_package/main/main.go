package main

import (
	"fmt"
	util "go_code/chapter06/function02_package/utils"
)

func main()  {

	var n1 float64 = 1.1
	var n2 float64 = 2.2
	var operator byte = '*'
	var result float64 = util.Calc(n1, n2, operator)

	fmt.Println("utils.go Num:", util.Num)
	fmt.Printf("utilsDupe.go PrintOK function:")
	util.PrintOk()
	fmt.Printf("Result:%.2f\n", result)
	fmt.Printf("Result:%.2f\n", util.Calc(4.5, 9.4, '-'))
}
