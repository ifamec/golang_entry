package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {
	// current time
	now := time.Now()
	fmt.Printf("%v, %T\n", now, now)

	fmt.Printf("year: %v month: %v day: %v \nhour: %v minute: %v second: %v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("current time:", dateStr)

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	fmt.Println(100 * time.Millisecond)

	// Print a number til 100, print 1 number per 0.1 second
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println(now.Unix(), now.UnixNano())

	var timeStart = time.Now().Unix()
	getExecutionTime()
	var timeEnd =time.Now().Unix()
	fmt.Println(timeEnd - timeStart)
}

func getExecutionTime()  {
	str := ""
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
}