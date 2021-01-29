package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "./src/go_code/chapter14/testFileWrite.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666) // READ WRITE APPEND
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Read
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	// Write 5 Sentence
	str01 := "Oh My God\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str01)
	}
	// Move content in buffer to disk
	writer.Flush()

}
