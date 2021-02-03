package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Age int
	Address string
}

func main()  {
	var personChan chan *Person = make(chan *Person, 10)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		var randomInt int = rand.Intn(100)
		personChan <- &Person{
			Name: "Name_" + string(rune(randomInt)),
			Age: randomInt,
			Address: "Address_" + string(rune(randomInt)),
		}
	}
	close(personChan)

	for val := range personChan {
		fmt.Println(*val)
	}

}