package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"mName"`
	Age      int    `json:"mAge"`
	Birthday string
	Salary   float64
	Skill    string
}

func deserializeStruct() {
	str := `{"mName":"AAA","mAge":20,"Birthday":"2000-01-01","Salary":50.5,"Skill":"Clime"}`

	// monster variable
	var m Monster
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("Deserialize Error:", err)
	}
	fmt.Println(m.Name, m.Age, m.Birthday, m.Salary, m.Skill)
}

func deserializeMap() {
	str := `{"Address":"Guangzhou","Age":19,"Name":"BBB"}`

	var m map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("Deserialize Error:", err)
	}
	fmt.Println(m)
}

func deserializeSlice() {
	str := `[{"Age":"18","Name":"CCC"},{"Age":17,"Name":"CCC"}]`

	var s []map[string]interface{}
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("Deserialize Error:", err)
	}
	fmt.Println(s)
}

func main() {
	deserializeStruct()

	deserializeMap()

	deserializeSlice()
}
