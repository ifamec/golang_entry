package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./src/go_code/chapter14/testFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("file = %v", file)

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}
