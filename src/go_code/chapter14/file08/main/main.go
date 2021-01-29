package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Args length:", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("%v\t%v\n", i, v)
	}
}
