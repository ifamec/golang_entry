package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var count int
	// to get a random number, set a seed for rand
	rand.Seed(time.Now().UnixNano()) // get a unix time second
	for {
		var num = rand.Intn(100) + 1
		count++
		if num == 99 {
			break
		}
	}
	fmt.Println("Times:", count)

	// break-label
	label2:
	for i:=0;i<4;i++{
		//label1: // set a label
		for j:=0;j<10;j++ {
			if j==2 {
				//break // jump out of this loop // 01010101
				//break label1 // 01010101
				break label2 // 01
			}
			fmt.Printf("%d", j)
		}
	}

	fmt.Println("")
	exercise()
}

func exercise()  {
	// sum of 1-100 print the number that make the sum > 20
	var sum int
	for i :=1; i<=100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println(i)
			break
		}
	}

	// login test
	for i := 2; i >= 0; i-- {
		var uname, passwd string
		fmt.Println("Input usrname passwd")
		fmt.Scanf("%s %s", &uname, &passwd)
		if uname == "admin" && passwd == "888" {
			fmt.Println("Success")
			break
		} else if (i == 0) {
			fmt.Printf("Error. You have no chance, please try next time\n")
		} else {
			fmt.Printf("Error. You have %d chance left\n", i)
		}
	}
}