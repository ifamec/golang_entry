package main

import (
	"fmt"
	"math"
)

func main() {
	// declare in if statement
	if n := 1; n > 0 {
		fmt.Printf("n declared in if statement: %v\n", n)
	}
	// if statement is a assignment
	//var x bool = true // COMPILE ERROR
	//if x = false {
	//	fmt.Println('ERROR')
	//}

	var age byte
	fmt.Printf("Input Age: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("S - You are an adult")
	}

	if age > 18 {
		fmt.Println("D - You are an adult")
	} else {
		fmt.Println("D - You are not an adult")
	}

	var score byte
	fmt.Printf("WHAI IS YOUR SCORE: ")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("A")
	} else if score > 80 && score <= 99 {
		fmt.Println("B")
	} else if score >= 60 && score <= 80 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}

	// get root
	var a, b, c float64
	fmt.Println("INPUT a b c separate with space")
	fmt.Scanf("%f %f %f\n", &a, &b, &c)

	var delta float64 = b*b - 4*a*c
	if delta > 0 {
		fmt.Println("TWO REAL ROOT:", (-b+math.Sqrt(delta))/2./a, (-b-math.Sqrt(delta))/2./a)
	} else if delta == 0 {
		fmt.Println("ONE REAL ROOT:", (-b+math.Sqrt(delta))/2./a)
	} else {
		fmt.Println("NO REAL ROOT")
	}

	// 3 condition
	var t byte
	var r float32
	var h bool
	fmt.Println("INPUT t(byte) r(float32) c(bool) separate with space")
	fmt.Scanf("%d %f %t\n", &t, &r, &h)

	if t > 180 && r > 1000 && h {
		fmt.Println("3")
	} else if t > 180 || r > 1000 || h {
		fmt.Println("1/2")
	} else {
		fmt.Println("0")
	}

	// NEST01
	var second float64
	var gender string
	fmt.Println("INPUT second(float64) gender(string M/F) separate with space")
	fmt.Scanf("%f %s\n", &second, &gender)

	if second < 8. {
		if gender == "M" {
			fmt.Println("Male")
		} else {
			fmt.Println("Female")
		}
	} else {
		fmt.Println("Fail")
	}

	// NEST 02
	var month byte
	var ageN byte
	fmt.Println("INPUT current month(byte) age(byte) separate with space")
	fmt.Scanf("%d %d\n", &month, &ageN)
	if month >=4 && month <= 10 {
		if ageN > 60 {
			fmt.Println("1/3")
		} else if ageN < 18 {
			fmt.Println("1/2")
		} else {
			fmt.Println("FULL")
		}
	} else if month > 0 && month <= 12 {
		if ageN <= 60 && ageN >= 18 {
			fmt.Println("FULL")
		} else {
			fmt.Println("1/2")
		}
	}


	fmt.Printf("\nEXERCISE 01:\n\n")
	exercise_1()
}

func exercise_1() {
	// declare 2 int32 variables, if sum >= 50 print "hello world"
	var a, b int32 = 25, 27
	if a+b >= 50 {
		fmt.Println("hello world")
	}
	// declare 2 float64 variables, if 1st > 10.0 and 2nd <20.0 print sum
	var c, d float64 = 20.21, 19.94
	if c > 10.0 && d < 20.0 {
		fmt.Println(c + d)
	}
	// declare 2 int32 variables, print if the sum can be divided by both 3 and 5
	var e, f int32 = 91, 29
	if (e+f)%15 == 0 {
		fmt.Println("YES IT CAN")
	}
	// is leap year 1) can be divided by 4 but not 100, 2) can be divided by 400
	var year int = 2020
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		fmt.Println("IS LEAP YEAR")
	}
}
