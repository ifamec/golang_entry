package model

import "fmt"

type person struct {
	Name string
	age int
	salary float64
}

func NewPerson(n string, a int, s float64) *person {
	return &person{n, a, s}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		(*p).age = age
	} else {
		fmt.Println("AGE INVALID")
	}
}
func (p *person) GetAge() int {
	return (*p).age
}
func (p *person) SetSalary(sal float64) {
	if sal > 3000 {
		(*p).salary = sal
	} else {
		fmt.Println("SALARY INVALID")
	}
}
func (p *person) GetSalary() float64 {
	return (*p).salary
}