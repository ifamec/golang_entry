package main

import (
	"fmt"
	"go_code/chapter15/test02/monster"
	"os"
)

func main()  {
	filePath := "./src/go_code/chapter15/test02/monster/store.txt"

	var m *monster.Monster = &monster.Monster{"AAA", 20, "Climbing"}
	(*m).Store(filePath)

	var mr *monster.Monster = &monster.Monster{}
	(*mr).ReStore(filePath)

	fmt.Println(*mr)

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Remove Test File Fails")
	}
}