package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main()  {

	var anyChan chan interface{} = make(chan interface{}, 3)

	anyChan <- 3
	anyChan <- "Tom"
	anyChan <- Cat{"Jerry", 3}

	<- anyChan
	<- anyChan

	cat := <- anyChan
	fmt.Printf("%T, %v\n", cat, cat)
	// fmt.Println(cat.Name) // Error: cat.Name undefined (type interface {} is interface with no methods)

	realCat := cat.(Cat) // USE TYPE ASSERTION
	fmt.Println(realCat.Name, realCat.Age)
}