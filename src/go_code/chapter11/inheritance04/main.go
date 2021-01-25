package main

import "fmt"

type A struct {
	Name string
	age int
}
type B struct {
	Name string
	score float64
}

type C struct {
	A
	B
}

type D struct {
	a A
	// Name string
}

type Goods struct {
	Name string
	Price float64
}
type Brand struct {
	Name string
	Address string
}
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods
	*Brand
}


type M struct {
	Name string
	Age int
}
type MM struct {
	M
	int
}

func main()  {
	var c = &C{}
	(*c).A.Name = "AAA"
	(*c).B.Name = "BBB"
	// IF C does not have Name field, but A and B have, NEED TO INDICATE ANONYMOUS STRUCT NAME TO USE
	// SAME TO METHODS
	// fmt.Println((*c).Name) // Ambiguous reference ERROR
	fmt.Println((*c).A.Name)
	fmt.Println((*c).B.Name)

	var d = &D{}
	// (*d).Name = "DDD"
	(*d).a.Name = "JACK"

	var tv = &TV{
		Goods{"Television", 1800.00},
		Brand{"Sony", "Tokyo Japan"},
	}
	fmt.Println((*tv).Goods.Name, (*tv).Brand.Name, (*tv).Price)
	var tv2 = &TV{
		Goods{
			Price:1800.00,
			Name: "Television",
		},
		Brand{
			Address: "Seoul Korea",
			Name: "LG",
		},
	}
	var tv3 = &TV2{
		&Goods{"Television", 1800.00},
		&Brand{"Sharp", "Tokyo Japan"},
	}
	fmt.Println(*tv, *tv2, *tv3)
	fmt.Println(*(*tv3).Goods, *(*tv3).Brand)


	var m0 = &MM{}
	(*m0).M.Name = "M0"
	(*m0).M.Age = 9
	(*m0).int = 10
	fmt.Println(*m0)
}