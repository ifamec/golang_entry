package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//ex1()
	//ex2()
	//ex3(100)
	//ex4("2020-01-20")
	//ex5()
	ex6()
}

func ex1() {
	// get yyyy, mm, dd, check if valid or not
	var y, m, d int

	var errorMsg = func(param string) {
		fmt.Println(param, "INPUT ERROR")
	}

	for {
		fmt.Println("Input Year:")
		fmt.Scanf("%d\n", &y)
		var FebLastDay int = func(year int) int {
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				return 29
			} else {
				return 28
			}
		}(y)
		fmt.Println("Input Month:")
		fmt.Scanf("%d\n", &m)
		if m > 12|| m < 1 {
			errorMsg("MONTH")
			continue
		}
		fmt.Println("Input Day:")
		fmt.Scanf("%d\n", &d)

		switch m {
		case 1, 3, 5, 7, 8, 10, 12:
			if d > 31 || d < 0 {
				errorMsg("DAY")
			} else {
				fmt.Printf("%d-%d-%d\n", y, m, d)
			}
		case 4, 6, 9, 11:
			if d > 30 || d < 0 {
				errorMsg("DAY")
			} else {
				fmt.Printf("%d-%d-%d\n", y, m, d)
			}
		case 2:
			if d > FebLastDay || d < 0 {
				errorMsg("DAY")
			} else {
				fmt.Printf("%d-%d-%d\n", y, m, d)
			}
		}
	}
}
func ex2() {
	// guess number
	rand.Seed(time.Now().UnixNano()) // get a unix time second
	// or use UnixNano()
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	var print = func (word string) {
		fmt.Println(word)
	}
	i := 1
	for {
		var v int
		fmt.Println("INPUT:")
		fmt.Scanf("%d\n", &v)
		if i < 11 {
			if v == num {
				switch {
				case i == 1: print("EXCELLENT")
				case i >= 2 && i <= 3: print("GOOD")
				case i >= 4 && i <= 9: print("SOSO")
				case i == 10: print("FINALLY")
				}
				break
			} else {
				i++
				continue
			}
		} else {
			print("OH NO")
			break
		}
	}
}
func ex3(n int) {
	// calc prime number print 5 each line, and sum
	var isPrime = func(m int) bool {
		if m == 2 || m==1 { return true }
		i := 2
		for ; i <= m; i++ {
			if m%i == 0 {
				break
			}
		}
		return m == i
	}
	var sum int
	var printIndicator int = 0
	for i:=1; i<=100; i++ {
		if printIndicator == 5 {
			fmt.Println("")
			printIndicator = 0
		}
		if isPrime(i) {
			sum += i
			fmt.Printf("%d\t", i)
			printIndicator++
		}
	}
	fmt.Printf("\nSUM: %d\n", sum)
}
func ex4(date string)  {
	ori, _ := time.Parse("2006-01-02", "1990-01-01")
	input, err := time.Parse("2006-01-02", date)
	var delta int = int(input.Sub(ori).Hours()/24) + 1
	if err != nil || delta < 1 {
		fmt.Println("INPUT ERROR")
		return
	}
	switch delta%5 {
	case 1,2,3: fmt.Println("FISHING")
	case 4,0:	fmt.Println("DRYING")
	}

}
func ex5() {
	var op byte
	for {
		fmt.Printf("1.+\n2.-\n3.*\n4./\n0.exit\nINPUT:")
		fmt.Scanf("%d\n", &op)
		switch op {
		case 1: fmt.Println("10+5=", 10 + 5)
		case 2: fmt.Println("10-5=", 10 - 5)
		case 3: fmt.Println("10*5=", 10 * 5)
		case 4: fmt.Println("10/5=", 10 / 5)
		case 0: fmt.Println("EXIT"); return
		default: fmt.Println("INVALID, TRY AGAIN")
		}
	}
}
func ex6() {
	for i := 'a'; i <= 'z'; i ++{
		fmt.Printf("%c\t", i)
	}
	fmt.Println("")
	for i := 'Z'; i >= 'A'; i --{
		fmt.Printf("%c\t", i)
	}
}