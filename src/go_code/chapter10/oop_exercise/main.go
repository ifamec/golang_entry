package main

import "fmt"

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}
func (s *Student) say() string {
	return fmt.Sprintf("Name: %v, Gender: %v, Age %v, Id: %v, Score: %v",
		(*s).name, (*s).gender, (*s).age, (*s).id, (*s).score)
}

type Dog struct {
	name string
	age int
	weight float64
}
func (d *Dog) say() string {
	return fmt.Sprintf("Name: %v, Age %v, Weight: %v",
		(*d).name, (*d).age, (*d).weight)
}

type Box struct {
	length float64
	width float64
	height float64
}
func (b *Box) getVolumn() float64 {
	return (*b).length * (*b).width * (*b).height
}

type Visitor struct {
	Name string
	Age int
}
func (v *Visitor) getTicketPrice() int {
	if (*v).Age > 18 {
		return 20
	}
	return 0
}

func main()  {
	var s1 Student = Student{"Tom", "M", 20, 1000001, 98.99}
	fmt.Println((&s1).say())

	var d1 Dog = Dog{"OOO", 20, 18.99}
	fmt.Println((&d1).say())

	var length, width, height float64 = 0, 0, 0
	fmt.Printf("INPUT L W H:")
	fmt.Scanf("%f %f %f\n", &length ,&width ,&height)
	var b1 Box = Box{length, width, height}
	fmt.Println(b1.getVolumn())


	var v1 Visitor
	for  {
		fmt.Printf("INPUT Name:")
		fmt.Scanf("%s\n", &v1.Name)
		if v1.Name == "n" {
			break
		}
		fmt.Printf("INPUT Age:")
		fmt.Scanf("%d\n", &v1.Age)
		fmt.Println(v1.Name, "is", v1.Age, "Years Old, Ticket Price is", v1.getTicketPrice())
	}

}