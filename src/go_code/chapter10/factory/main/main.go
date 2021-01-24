package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main()  {
	//s1 := model.Student{"Tom", 97.99}
	s1 := model.NewStudent("Tom", 97.99) // s1 is a pointer to the struct variable
	fmt.Println(*s1)
	fmt.Println((*s1).Name, (*s1).GetScore())
	s1.SetScore(99.99)
	fmt.Println((*s1).Name, (*s1).GetScore())
}