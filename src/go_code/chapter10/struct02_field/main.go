package main

import "fmt"

type Reference struct {
	Name string
	Age int
	Score [6]float64
	ptr *int
	slice []int
	map1 map[string]string
}

type Monster struct {
	Name string
	Age int
}

func main() {
	var ref Reference

	fmt.Println(ref)
	if ref.ptr == nil{
		fmt.Println("ref.ptr is nil by default")
	}
	if ref.slice == nil{
		fmt.Println("ref.slice is nil by default")
	}
	ref.slice = make([]int, 10)
	fmt.Println(ref.slice)
	if ref.map1 == nil{
		fmt.Println("ref.map1 is nil by default")
	}
	ref.map1 = make(map[string]string)
	if ref.map1 != nil{
		fmt.Println("after make, ref.map1 is exist")
	}
	fmt.Println(ref.map1)

	m1 := Monster{"M1", 10}
	m2 := m1
	fmt.Printf("%p\t%p\n", &m1, &m2)
	fmt.Println(m1, m2)
	m2.Name = "M2"
	fmt.Println(m1, m2)
}
