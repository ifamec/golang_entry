package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hay", i+1)
	}

	j := 0
	for j < 10 {
		fmt.Println("Hey", j+1)
		j++
	}

	k := 0
	for { // use with break == for ;; {
		if k < 10 {
			fmt.Println("Yos", k+1)
		} else {
			break
		}
		k++
	}

	var str string = "Hello, World广州"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println("")

	var str2 = []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}

	fmt.Println("")
	for index, val := range str{
		fmt.Printf("%v, %c \t", index, val)
	}

	fmt.Println("EXERCISE")
	exercise()

	// while
	while := 0
	for {
		if while > 9 {
			break
		}
		fmt.Println("while,", while)
		while++
	}

	// do while
	doWhile := 0
	for {
		fmt.Println("do_while,", doWhile)
		doWhile++
		if doWhile > 9 {
			break
		}
	}

	exercise2()
}

func exercise()  {
	// print 9x between 1-100 and sum
	var sum, count int
	for i := 1; i < 101; i++ {
		if i%9 == 0 {
			fmt.Println(i)
			sum += i
			count ++
		}
	}
	fmt.Printf("sum:%d count:%d\n", sum, count)

	// print
	for j := 0; j < 7; j++ {
		fmt.Printf("%d  +  %d  =  6\n", j, 6 - j)
	}

}

func exercise2()  {
	fmt.Println("---------- Nested Loop Example ----------")
	ex2_1()
	ex2_1_imp()
	ex2_2()
	ex2_3(6)
	ex2_3_i(6)
	ex2_3_while(5)
	ex2_4(10)
}

func ex2_1() {
	// 3 classes each class 5 stu, get each class avg and all avg
	fmt.Println("---------- Nested Loop Example 1 ----------")
	var total float64
	for i:=0;i<3;i++ {
		var sum float64
		for j:=0;j<5;j++ {
			var score float64
			fmt.Println("input score")
			fmt.Scanf("%f\n", &score)
			sum += score
			total += score
		}
		fmt.Printf("avg of class %d is %f\n", i, sum/5.)
	}
	fmt.Printf("avg of total is %f\n", total/15.)
}
func ex2_1_imp() {
	fmt.Println("---------- Nested Loop Example 1 imp ----------")
	// 3 classes each class 5 stu, get each class avg and all avg
	var classNum, stuNum int = 3, 5
	var total float64
	for i := 0; i < classNum; i++ {
		var sum float64
		for j := 0; j < stuNum; j++ {
			var score float64
			fmt.Println("input score")
			fmt.Scanf("%f\n", &score)
			sum += score
			total += score
		}
		fmt.Printf("avg of class %d is %f\n", i, sum/float64(stuNum))
	}
	fmt.Printf("avg of total is %f\n", total/float64(classNum*stuNum))
}
func ex2_2()  {
	fmt.Println("---------- Nested Loop Example 2 ----------")
	// 3 classes each class 5 stu, get each class avg and all avg
	var passCount int = 0
	var classNum, stuNum int = 3, 5
	var total float64
	for i := 0; i < classNum; i++ {
		var sum float64
		for j := 0; j < stuNum; j++ {
			var score float64
			fmt.Println("input score")
			fmt.Scanf("%f\n", &score)
			sum += score
			total += score
			if score >= 60 {
				passCount ++
			}
		}
		fmt.Printf("avg of class %d is %f\n", i, sum/float64(stuNum))
	}
	fmt.Printf("avg of total is %f\n", total/float64(classNum*stuNum))
	fmt.Printf("pass count is %d\n", passCount)
}
func ex2_3(layer int)  {
	// right pyramid
	for i := 1; i <= layer; i ++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	// full pyramid
	for i := 1; i <= layer; i ++ {
		for j := 1; j <= layer - i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	// empty pyramid
	for i := 1; i <= layer; i ++ {
		for j := 1; j <= layer - i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == layer {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
func ex2_3_i(layer int)  {
	// right pyramid
	for i := 1; i <= layer*2-1; i ++ {
		if i <= layer {
			for j := 1; j <= i; j++ {
				fmt.Printf("*")
			}
		} else {
			for j := 1; j <= layer * 2 - i; j++ {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
	// full pyramid
	for i := 1; i <= layer*2-1; i ++ {

		if i <= layer {
			for j := 1; j <= layer - i; j++ {
				fmt.Printf(" ")
			}
			for k := 1; k <= 2*i- 1; k++ {
				fmt.Printf("*")
			}
		} else {
			for j := 1; j <= i - layer; j++ {
				fmt.Printf(" ")
			}
			for k := 1; k <= 4*layer - 1 - 2*i; k++ { //7 1 9 - 8 2 7 - 9 3 5
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
	// empty pyramid
	for i := 1; i <= layer*2-1; i ++ {

		if i <= layer {
			for j := 1; j <= layer - i; j++ {
				fmt.Printf(" ")
			}
			for k := 1; k <= 2*i- 1; k++ {
				if k == 1 || k == 2*i-1 {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
		} else {
			for j := 1; j <= i - layer; j++ {
				fmt.Printf(" ")
			}
			for k := 1; k <= 4*layer - 1 - 2*i; k++ {
				if k == 1 || k == 4*layer - 1 - 2*i {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}
}
func ex2_3_while(layer int)  {
	// right pyramid
	var i int

	i = 1
	for {
		if i > layer {
			break
		}
		var j int = 1
		for {
			if j > i {
				break
			}
			fmt.Printf("*")
			j++
		}
		fmt.Printf("\n")
		i++
	}
	// full pyramid
	i = 1
	for {
		if i > layer {
			break
		}
		var j, k int = 1, 1
		for {
			if j > layer - i {
				break
			}
			fmt.Printf(" ")
			j++
		}
		for {
			if k > 2*i-1 {
				break
			}
			fmt.Printf("*")
			k++
		}
		fmt.Printf("\n")
		i++
	}
	// empty pyramid
	i = 1
	for {
		if i > layer {
			break
		}
		j, k :=1, 1
		for {
			if j > layer - i {
				break
			}
			fmt.Printf(" ")
			j++
		}
		fmt.Printf("*")
		for {
			if k > 2*i-3 {
				break
			}
			if i == layer {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
			k++
		}
		if k > 1 {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
		i++
	}
}
func ex2_4(number int)  {
	fmt.Println("---------- Times Table ----------")

	for i := 1; i <=number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println("")
	}
}