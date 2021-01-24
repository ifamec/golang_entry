package main

import "fmt"

type MethodUtils struct {

}

func (mu *MethodUtils) printRecStatic() {
	for i := 0; i < 10; i ++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func (mu *MethodUtils) printRect(m int, n int) {
	for i := 0; i < m; i ++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func (mu *MethodUtils) getRectArea(m float64, n float64) float64 {
	return m * n
}
func (mu *MethodUtils) oddOrEven(m int) {
	if m % 2 == 0 {
		fmt.Println(m ,"IS EVEN")
	} else {
		fmt.Println(m ,"IS ODD")
	}
}
func (mu *MethodUtils) printRectCustom(m int, n int, key string) {
	for i := 0; i < m; i ++ {
		for j := 0; j < n; j++ {
			fmt.Printf(key)
		}
		fmt.Println()
	}
}
func (mu *MethodUtils) productTable(m int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", i, j, i*j)
		}
		fmt.Println()
	}
}
func (mu *MethodUtils) transposition(arr *[3][3]int) {
	for _, v := range *arr {
		fmt.Println(v)
	}
	// transposition
	for i := 0; i < len(*arr); i++ {
		for j := i; j < len((*arr)[0]); j++ {
			if i != j {
				(*arr)[i][j], (*arr)[j][i] = (*arr)[j][i], (*arr)[i][j]
			}
		}
	}
	for _, v := range *arr {
		fmt.Println(v)
	}
}

type Calculator struct {
	Num1 int
	Num2 int
}
func (cal *Calculator) getSum() int {
	return (*cal).Num1 + (*cal).Num2
}
func (cal *Calculator) getDiff() int {
	return (*cal).Num1 - (*cal).Num2
}
func (cal *Calculator) getProduct() int {
	return (*cal).Num1 * (*cal).Num2
}
func (cal *Calculator) getQuotient() float64 {
	if (*cal).Num2 == 0 {
		panic("0 cannot be divided")
	}
	return float64((*cal).Num1) / float64((*cal).Num2)
}
func (cal *Calculator) Calc(op byte) (rtnval float64) {

	if op == '/' && (*cal).Num2 == 0 {
		panic("0 cannot be divided")
	}
	n1, n2 := float64((*cal).Num1), float64((*cal).Num2)
	switch op {
	case '+': rtnval = n1 + n2
	case '-': rtnval = n1 - n2
	case '*': rtnval = n1 * n2
	case '/': rtnval = n1 / n2
	default:
		fmt.Println("Please Provide Correct Operator")
	}
	return
}

func main()  {
	var mu MethodUtils = MethodUtils{}
	(&mu).printRecStatic()
	fmt.Println()

	(&mu).printRect(3, 10)
	fmt.Println()

	area := (&mu).getRectArea(3, 10.5)
	fmt.Println("Area is", area)
	fmt.Println()

	(&mu).oddOrEven(5)
	fmt.Println()

	(&mu).printRectCustom(3, 6, "(√)")
	fmt.Println()

	fmt.Println()
	var calc Calculator = Calculator{9,5}
	fmt.Println("getSum:", calc.getSum())
	fmt.Println("getDiff:", calc.getDiff())
	fmt.Println("getProduct:", calc.getProduct())
	fmt.Println("getQuotient:", calc.getQuotient())
	fmt.Println()
	fmt.Println("getSum:", calc.Calc('+'))
	fmt.Println("getDiff:", calc.Calc('-'))
	fmt.Println("getProduct:", calc.Calc('*'))
	fmt.Println("getQuotient:", calc.Calc('/'))


	var matrix [3][3]int = [3][3]int{
		{1,2,3},{4,5,6},{7,8,9},
	}
	(&mu).productTable(9)
	(&mu).transposition(&matrix)
	fmt.Println(matrix)

}