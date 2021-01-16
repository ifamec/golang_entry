package main

import "fmt" // format IO func

func main()  {
	// excape char example
	fmt.Println("Tom\tCat")
	fmt.Println("Jerry\nMouse")
	fmt.Println("C:\\Users\\user\\go\\src")
	fmt.Println("Tom said: \"Hi\".")
	fmt.Println("TomCat\rCat") // Replace from beginning

	// Exercise
	fmt.Println("Name\tAge\tHometown\tCity\nTom\t23\tZhejiang\tGuangzhou")

	// limit 80 chars in one line
	fmt.Println("aaaaaaaaaaaaaaaaaaaa",
		"bbbbbbbbbbbbbbbbbbbb",
		"cccccccccccccccccccb",
		"dddddddddddddddddddb")

	// Exercise
	fmt.Println("\t*\t\t\t\t\t\t\t\t\t\t\t\t*\t\n" +
		"*\t\t\t*\t\tI love Golang\t\t\t*\t\t\t*\n" +
		"\t*\t\t\t\t\t\t\t\t\t\t\t\t*\n" +
		"\t\t*\t\t\t\t\t\t\t\t\t\t*\n" +
		"\t\t\t*\t\t\t\t\t\t\t\t*\n" +
		"\t\t\t\t*\t\t\t\t\t\t*\n" +
		"\t\t\t\t\t*\t\t\t\t*\n" +
		"\t\t\t\t\t\t*\t\t*\n" +
		"\t\t\t\t\t\t\t*")
}

