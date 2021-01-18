package main

import "fmt"

func main()  {

	var n1 float64 = 1.1
	var n2 float64 = 2.2
	var operator byte = '*'
	var result float64 = calc(n1, n2, operator)

	//switch operator {
	//case '+': result = n1 + n2
	//case '-': result = n1 - n2
	//case '*': result = n1 * n2
	//case '/': result = n1 / n2
	//default: fmt.Println("Operator Error")
	//}

	fmt.Println("Result:", result)
	fmt.Println("Result:", calc(4.5, 9.4, '-'))
}

func calc(n1 float64, n2 float64, operator byte) (result float64){

	switch operator {
	case '+': result = n1 + n2
	case '-': result = n1 - n2
	case '*': result = n1 * n2
	case '/': result = n1 / n2
	default: fmt.Println("Operator Error")
	}
	return result
}