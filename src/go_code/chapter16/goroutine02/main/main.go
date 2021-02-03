package main

import (
	"fmt"
	"runtime"
)

func main() {
	// check logical cpus
	num := runtime.NumCPU()
	fmt.Println(num)

	// set go maximum process number
	runtime.GOMAXPROCS(num - 1)

}
