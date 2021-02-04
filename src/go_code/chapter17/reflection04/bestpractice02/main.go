package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Use reflection traverse struct fields
// call methods in struct to get the struct tag value

type Monster struct {
	Name   string `json:"name"`
	Age    int    `json:"monster_age"`
	Score  float32
	Gender string
}

func (s Monster) Print() { // method 1
	fmt.Println("--- start ---\n", s, "\n--- end ---")
}

func TestStruct(a interface{}) {
	rType := reflect.TypeOf(a) 		// get reflect.Type
	rValue := reflect.ValueOf(a)	// get reflect.Value
	rKind := rValue.Kind()			// get Kind of a

	if rKind != reflect.Ptr && rValue.Elem().Kind() == reflect.Struct {	// Make sure a is a struct
		fmt.Println("Expect struct")
		return
	}

	num := rValue.Elem().NumField()		// Get Number of fields
	rValue.Elem().Field(0).SetString("M3")
	for i := 0; i < num; i++ {
		fmt.Println("Field", i, " Kind is", rValue.Elem().Field(i).Kind())
	}
	fmt.Printf("struct has %d fields\n", num)

	tagValue := rType.Elem().Field(0).Tag.Get("json")
	fmt.Println("Field 0", " Tag is", tagValue)


	numOfMethods := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethods)
	rValue.Method(0).Call(nil) // Print

}

func main() {
	var m Monster = Monster{
		Name: "M1",
		Age: 18,
		Score: 98.99,
	}
	result, _ := json.Marshal(m)
	fmt.Println("JSON:", result)

	TestStruct(&m)
	fmt.Println(m)
}
