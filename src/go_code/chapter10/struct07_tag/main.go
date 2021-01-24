package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name string 	`json:"name"`
	Age int			`json:"age"`
	Color string	`json:"color"`
}

func main() {
	var cat1 Cat = Cat{"Oolong", 4, "Orange"}
	str, _ := json.Marshal(cat1)
	fmt.Println(string(str))
	// if change filed name to all lowercase, then json cannot use the struct here (scope)
}
