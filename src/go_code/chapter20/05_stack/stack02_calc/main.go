package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	maxSize int     // maxSize
	Top     int     // stack top
	Array   [20]int // mock stack
}

func (s *Stack) Push(n int) (err error) {
	if (*s).Top == (*s).maxSize-1 {
		err = errors.New("stack full")
		return
	}
	(*s).Top++
	(*s).Array[(*s).Top] = n
	return
}

func (s *Stack) Pop() (n int, err error) {
	if (*s).Top == -1 {
		err = errors.New("stack empty")
		return
	}
	n = (*s).Array[(*s).Top]
	(*s).Top--
	return
}

func (s *Stack) List() (err error) {
	if (*s).Top == -1 {
		err = errors.New("stack empty")
		return
	}
	fmt.Println("List Stack")
	for i := (*s).Top; i >= 0; i-- {
		fmt.Println(i, (*s).Array[i])
	}
	return
}

func (s *Stack) IsOperator(val int) bool {
	return val == 42 || val == 43 || val == 45 || val == 47
}

func (s *Stack) Calculate(n1, n2, op int) (rtnval int) {
	switch op {
	case 42:
		rtnval = n2 * n1
	case 43:
		rtnval = n2 + n1
	case 45:
		rtnval = n2 - n1
	case 47:
		rtnval = n2 / n1
	default:
		fmt.Println("Operator Error")
	}
	return
}

func (s *Stack) OperatorPriority(op int) (rtnval int) {
	if op == 42 || op == 47 {
		return 2
	} else if op == 43 || op == 45 {
		return 1
	}
	return 0
}

func main() {
	numberStack := &Stack{
		maxSize: 20,
		Top:     -1,
	}
	operatorStack := &Stack{
		maxSize: 20,
		Top:     -1,
	}
	exp := "30+3*60/10-4*2-2"
	index := 0

	var num1, num2, op, rst int
	var keepNumber string

	for {
		// handle multiple numbers
		ch := exp[index : index+1]
		value := int([]byte(ch)[0])
		if operatorStack.IsOperator(value) {
			if operatorStack.Top == -1 { // empty stack
				_ = operatorStack.Push(value)
			} else {
				if operatorStack.OperatorPriority(operatorStack.Array[operatorStack.Top]) >= operatorStack.OperatorPriority(value) {
					num1, _ = numberStack.Pop()
					num2, _ = numberStack.Pop()
					op, _ = operatorStack.Pop()
					rst = operatorStack.Calculate(num1, num2, op)
					_ = numberStack.Push(rst)
					_ = operatorStack.Push(value)
				} else {
					_ = operatorStack.Push(value)
				}

			}
		} else { // number
			// check next val is operator or not, if EOF, keepnumber to number
			keepNumber += ch
			if index == len(exp) - 1 {
				v, _ := strconv.Atoi(keepNumber)
				_ = numberStack.Push(v)
			} else {
				if operatorStack.IsOperator(int( []byte( exp[index+1:index+2] ) [0])) {
					v, _ := strconv.Atoi(keepNumber)
					_ = numberStack.Push(v)
					keepNumber = ""
				}
			}

		}
		// continue
		if index+1 == len(exp) {
			break
		}
		index++
	}

	// calculate
	for {
		if operatorStack.Top == -1 {
			break
		}
		num1, _ = numberStack.Pop()
		num2, _ = numberStack.Pop()
		op, _ = operatorStack.Pop()
		rst = operatorStack.Calculate(num1, num2, op)
		_ = numberStack.Push(rst)
	}

	// if everything is good
	result, _ := numberStack.Pop()
	fmt.Println(exp, "=" , result)

}
