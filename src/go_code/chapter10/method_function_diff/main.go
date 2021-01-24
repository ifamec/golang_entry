package main

import "fmt"

type Person struct {
	Name string
}

func test01(p Person)  {
	fmt.Println("Function", p.Name)
}
func test02(p *Person)  {
	fmt.Println("Function *", p.Name)
}

func (p Person) test03() {
	p.Name = "JERRY"
	fmt.Println("Method", p.Name)
}
func (p *Person) test04() {
	(*p).Name = "Mary"
	fmt.Println("Method *", (*p).Name)
}


func main()  {
	var p Person = Person{
		Name: "TOM",
	}
	//test01(&p) // ERROR CANNOT PASS POINTER
	test01(p)
	test02(&p) // MUST PASS POINTER

	fmt.Println()

	fmt.Println("MAIN", p.Name)
	p.test03()
	fmt.Println("MAIN", p.Name)

	fmt.Println()

	fmt.Println("MAIN", p.Name)
	(&p).test03() // call using pointer, but it's a pass by value
	fmt.Println("MAIN", p.Name)

	fmt.Println()

	fmt.Println("MAIN", p.Name)
	(&p).test04()
	fmt.Println("MAIN", p.Name)

	fmt.Println()

	// revert
	p.Name = "TOM Revert"
	fmt.Println("MAIN", p.Name)
	p.test04()
	fmt.Println("MAIN", p.Name)
}