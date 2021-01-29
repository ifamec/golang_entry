package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Count struct {
	Alphabet int
	Number   int
	Space    int
	Other    int
}

func main() {
	var count Count = Count{}
	path := "./src/go_code/chapter14/file07/main/Count.txt"
	// readline, iterate, save to struct
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open Error:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// iterate str
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.Alphabet++
			case v == ' ' || v == '\t':
				count.Space++
			case v >= '0' && v <= '9':
				count.Number++
			default:
				count.Other++
			}
		}
	}
	fmt.Println(count)

}
