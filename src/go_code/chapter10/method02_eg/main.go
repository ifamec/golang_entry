package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
}

func (c Person) speak() {
	fmt.Println(c.Name, "is a good person")
}
func (c Person) calc() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("Calc:", sum)
}
func (c Person) cal(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("Cal:", sum)
}
func (c Person) getSum(n int, m int) int {
	return n + m
}

type Circle struct {
	Radius float64
}

func (c Circle) getArea() float64 {
	return math.Pi * c.Radius * c.Radius
}
// To be more effective method bind to struct pointer
func (c *Circle) getArea2() float64 {
	fmt.Printf("getArea c address: %p\n", c)
	//return math.Pi * (*c).Radius * (*c).Radius // Standard Implementation
	return math.Pi * c.Radius * c.Radius // Compiler will handle this
}

func main()  {
	var p0 Person = Person{"Tom"}
	p0.speak()
	p0.calc()
	p0.cal(10)
	fmt.Println("getSum:", p0.getSum(1,1))

	var c Circle = Circle{3}
	fmt.Println("Area of the circle is:", c.getArea())
	var c2 Circle = Circle{5}
	fmt.Printf("main c address: %p\n", &c2)
	fmt.Println("Area of the circle is:", (&c2).getArea2()) // Standard Implementation
	c2.Radius = 10
	fmt.Println("Area of the circle is:", c2.getArea2()) // Compiler will handle this
}