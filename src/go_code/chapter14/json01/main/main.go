package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"mName"`
	Age      int	`json:"mAge"`
	Birthday string
	Salary   float64
	Skill    string
}

func serializeStruct() {
	monster := Monster{
		Name:     "AAA",
		Age:      20,
		Birthday: "2000-01-01",
		Salary:   50.5,
		Skill:    "Clime",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("Serialize Error:", err)
	}
	fmt.Printf("%s\n", data)
}

func serializeMap() {
	var m map[string]interface{} = make(map[string]interface{})
	m["Name"] = "BBB"
	m["Age"] = 19
	m["Address"] = "Guangzhou"
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("Serialize Error:", err)
	}
	fmt.Printf("%s\n", data)
}

func serializeSlice() {
	var s []map[string]interface{}
	var m1 = make(map[string]interface{})
	m1["Name"] = "CCC"
	m1["Age"] = "18"

	var m2 = map[string]interface{}{
		"Name": "CCC",
		"Age": 17,
	}
	s = append(s, m1)
	s = append(s, m2)

	data, err := json.Marshal(&s)
	if err != nil {
		fmt.Println("Serialize Error:", err)
	}
	fmt.Printf("%s\n", data)
}
func serializeFloat64 () {
	var f float64 = 3.1415
	data, err := json.Marshal(&f)
	if err != nil {
		fmt.Println("Serialize Error:", err)
	}
	fmt.Printf("%s\n", data)
}
func main() {
	// struct
	serializeStruct()
	// map
	serializeMap()
	// slice
	serializeSlice()
	// float
	serializeFloat64()
}
