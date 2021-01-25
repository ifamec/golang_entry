package main

import "fmt"

type Student struct{}

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		i++
		switch x.(type) {
		case bool:
			fmt.Println(i, x, "bool")
		case float64:
			fmt.Println(i, x, "float64")
		case int, int64:
			fmt.Println(i, x, "int")
		case nil:
			fmt.Println(i, x, "nil")
		case string:
			fmt.Println(i, x, "string")
		case Student:
			fmt.Println(i, x, "Student")
		case *Student:
			fmt.Println(i, x, "*Student")
		default:
			fmt.Println(i, x, "DEFAULT")
		}
	}
}

func main() {
	var stu Student
	var stuptr *Student = &stu
	TypeJudge(stu, stuptr)
	// TypeJudge('a', 10, 1.22, "OHMYGOD", [3]int{1, 2, 3})
}
