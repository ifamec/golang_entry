package main

import (
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
func (s Monster) GetSum(n1, n2 int) int { // method 0
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, gender string) { // method 2
	s.Name = name
	s.Age = age
	s.Score = score
	s.Gender = gender
}

func TestStruct(a interface{}) {
	rType := reflect.TypeOf(a) 		// get reflect.Type
	rValue := reflect.ValueOf(a)	// get reflect.Value
	rKind := rValue.Kind()			// get Kind of a

	if rKind != reflect.Struct {	// Make sure a is a struct
		fmt.Println("Expect struct")
		return
	}

	num := rValue.NumField()		// Get Number of fields
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Println("Field", i, " Value is", rValue.Field(i))
		tagValue := rType.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Println("Field", i, " Tag is", tagValue)
		}
	}

	numOfMethods := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethods)
	rValue.Method(1).Call(nil) // Print

	// Method Index: sorted by method name ASCII

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rValue.Method(0).Call(params) // Call GetSum, []reflect.Value
	fmt.Println("Res:", res)
	fmt.Println("Res:", res[0].Int())
}

func main() {
	var m Monster = Monster{
		Name: "M1",
		Age: 10,
		Score: 99.99,
	}
	TestStruct(m)
}
