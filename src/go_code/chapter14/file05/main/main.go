package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileFromPath := "./src/go_code/chapter14/file05/testFileFrom.txt"
	fileToPath := "./src/go_code/chapter14/file05/testFileTo.txt"
	data, err := ioutil.ReadFile(fileFromPath)
	if err != nil {
		fmt.Println("Read Error:", err)
		return
	}
	err = ioutil.WriteFile(fileToPath, data, 0666)
	if err != nil {
		fmt.Println("Write Error:", err)
		return
	}
}
