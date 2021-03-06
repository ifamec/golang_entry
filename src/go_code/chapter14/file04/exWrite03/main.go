package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./src/go_code/chapter14/testFileWrite.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666) // APPEND
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Write 5 Sentence
	str01 := "ABC SEATTLE\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str01)
	}
	// Move content in buffer to disk
	writer.Flush()

}
