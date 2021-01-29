package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	path := "./src/go_code/chapter14/file06/main/"
	// copy file
	written, err := CopyFile(path + "gopher_copy.png", path + "gopher_src.png")
	if err != nil {
		fmt.Println("Copy Error:", err)
	}
	fmt.Println(written, "bytes written.")

}

func CopyFile(dst string, src string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Open Error:", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Write Error:", err)
		return
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}