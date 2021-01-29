package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./src/go_code/chapter14/testFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close() // Close before function ends

	// default buffer is 4096 bytes
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // end by each line
		if err == io.EOF { break }
		fmt.Print(str)
	}
	fmt.Println("=== File Read End")
}
