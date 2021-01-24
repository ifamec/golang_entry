package main

import "fmt"

type Integer int

func (i Integer) print() {
	fmt.Println(i)
}
func (i *Integer) increase() {
	*i++
}

type Cat struct {
	Name string
	Age int
}

func (c *Cat) String() string {
	return fmt.Sprintf("Name: %v Age: %v", (*c).Name, (*c).Age)
}
func main()  {
	var i Integer = Integer(10)
	i.print()
	i.increase()
	i.print()

	var c Cat = Cat{"Oolong", 4}
	fmt.Println(c, &c)
}