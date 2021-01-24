package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {
	var s1 = Stu{"AA", 10}
	s2 := Stu{"BB", 20}
	var s3 Stu = Stu{
		Name: "CC",
		Age:  30,
	}
	var s4 Stu = Stu{
		Age:  40,
		Name: "DD",
	}
	fmt.Println(s1, s2, s3, s4)

	var t1 = &Stu{"EE", 50} // t1 -> address -> struct variable
	t2 := &Stu{"FF", 60}
	fmt.Println(t1, t2)

	var u1 = &Stu{
		Name: "GG",
		Age: 70,
	}
	u2 := &Stu{
		Age: 70,
		Name: "GG",
	}
	fmt.Println(*u1, *u2)

}
