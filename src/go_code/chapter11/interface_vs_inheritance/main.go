package main

import "fmt"

type Monkey struct {
	Name string
}

func (m *Monkey) Climbing() {
	fmt.Println((*m).Name, "Can Climbing")
}

type Bird interface {
	Flying()
}
type Fish interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey
}
func (lm *LittleMonkey) Flying() {
	fmt.Println((*lm).Name, "Can Flying")
}
func (lm *LittleMonkey) Swimming() {
	fmt.Println((*lm).Name, "Can Swimming")
}

func main()  {
	var littleMonkey LittleMonkey = LittleMonkey{Monkey{
		Name: "WK",
	}}
	littleMonkey.Climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()

}