package main

import "fmt"

func main() {
	var monster []map[string]string = make([]map[string]string, 1)
	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "A1"
		monster[0]["age"] = "500"
	}
	fmt.Println(monster, len(monster))
	// append
	monster = append(monster, map[string]string{"name": "A2", "age": "400"})
	fmt.Println(monster, len(monster))
	monster = append(monster, map[string]string{"name": "A3", "age": "300"})
	fmt.Println(monster, len(monster))
}
