package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// io/util ReadFile read whole file
	file := "./src/go_code/chapter14/testFile.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Printf("FIle Content: \n%s", content)
}
