package main

import "fmt"

func main()  {

	var address string = "Guangzhou Guangdong"

	fmt.Println(address)

	var str = "Hello"
	//str[0] = 'A' // Cannot modify the content of a string
	fmt.Println(str)

	var str2 = "\"abc\"\n"
	fmt.Println(str2)

	str3 := `
	var str2 = "\"abc\"\n"
	fmt.Println(str2)
	`
	fmt.Println(str3)

	str4 := "iPad " + "Pro"
	fmt.Println(str4)

	//str5 := "aaaaaaaaaaaaaaaaaaaa" // Error
	//	+ "bbbbbbbbbbbbbbbbbbbb"
	//	+ "cccccccccccccccccccc"
	//	+ "dddddddddddddddddddd"
	str5 := "aaaaaaaaaaaaaaaaaaaa" +
		"bbbbbbbbbbbbbbbbbbbb" +
		"cccccccccccccccccccc" +
		"dddddddddddddddddddd"
	fmt.Println(str5)


}