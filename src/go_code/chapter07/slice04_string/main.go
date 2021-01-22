package main

import "fmt"

func main()  {
	var str string = "Hello, World"

	slice := str[7:]
	fmt.Println(slice)

	// str => 'Zello, World'
	arr := []byte(str)
	arr[0] = 'Z'
	str = string(arr)
	fmt.Println(str)

	// use UTF8 ? some UTF8 char is 3 bytes
	arr1 := []rune(str)
	arr1[0] = '粤'
	str = string(arr1)
	fmt.Println(str)

}