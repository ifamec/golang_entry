package main

import (
	"errors"
	"fmt"
)

func test()  {
	// use defer +recover to capture and handle error
	defer func() {
		//err := recover() // builtin fn to capture error
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	rst := num1 / num2
	fmt.Println(rst)
}

func readConf(name string) error { // read `init.conf` if file name error, throw custom error
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("READ FILE ERROR")
	}
}
func test2()  {
	err := readConf("config.init")
	if err != nil {
		panic(err) // throw error and end program
	}
	fmt.Println("Continue Executing")
}
func main()  {
	test()

	fmt.Println("Error Handling")

	test2()
	fmt.Println("Continue Main")
}

