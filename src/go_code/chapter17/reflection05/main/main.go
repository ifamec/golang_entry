package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}
func (c *Cal) GetDiff(name string) {
	fmt.Printf("%s did GetDiff, %v - %v = %v\n", name, (*c).Num1, (*c).Num2, (*c).Num1 - (*c).Num2)
}

func ReflectCal(r interface{}) {
	rValue := reflect.ValueOf(r)

	s := rValue.Elem() // struct variable
	fNum := s.NumField()
	for i := 0; i< fNum; i++ {
		fmt.Println("Field", i,"Value is", s.Field(i))
	}

	var callSlide []reflect.Value
	callSlide = append(callSlide, reflect.ValueOf("Tom"))
	rValue.Method(0).Call(callSlide)

}

func main() {
	c := &Cal{100, 50}
	ReflectCal(c)

}
