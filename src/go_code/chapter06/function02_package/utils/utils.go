package utils

import "fmt"

var Num int = 1
func Calc(n1 float64, n2 float64, operator byte) (result float64){

	switch operator {
	case '+': result = n1 + n2
	case '-': result = n1 - n2
	case '*': result = n1 * n2
	case '/': result = n1 / n2
	default: fmt.Println("Operator Error")
	}
	return result
}