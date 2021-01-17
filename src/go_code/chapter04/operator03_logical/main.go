package main

import "fmt"

func test() bool {
	fmt.Println("Test...")
	return true
}
func main()  {
	var age int = 40

	// &&
	if age > 30 && age < 50 {
		fmt.Println("OK 40 AND")
	}
	if age > 30 && age < 40 {
		fmt.Println("OK 39 AND")
	}

	// ||
	if age > 30 || age < 40 {
		fmt.Println("OK 40 OR")
	}
	if age > 30 || age < 40 {
		fmt.Println("OK 39 OR")
	}

	// !
	if age > 30 {
		fmt.Println("OK 40 NOT")
	}
	if ! (age > 30) {
		fmt.Println("OK 39 NOT")
	}

	var i int = 10
	if i > 9 && test() {
		fmt.Println("OK &&")
	}
	if i < 9 && test() { // break before &&, test() not executed
		fmt.Println("NOK &&")
	}


	if i < 9 || test() { // "Test... \n OK ||"
		fmt.Println("OK ||")
	}
	if i > 9 || test() { // "NOK ||" only, test() not executed
		fmt.Println("NOK ||")
	}

}
