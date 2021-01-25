package main

import "fmt"

type Primary struct {
	Name  string
	Age   int
	Score float64
}

func (p *Primary) ShowInfo() {
	fmt.Println("Name:", (*p).Name, "Age:", (*p).Age, "Score", (*p).Score)
}
func (p *Primary) SetScore(s float64) {
	(*p).Score = s
}
func (p *Primary) Testing() {
	fmt.Println("Primary Testing")
}


type College struct {
	Name  string
	Age   int
	Score float64
}

func (p *College) ShowInfo() {
	fmt.Println("Name:", (*p).Name, "Age:", (*p).Age, "Score", (*p).Score)
}
func (p *College) SetScore(s float64) {
	(*p).Score = s
}
func (p *College) Testing() {
	fmt.Println("Primary Testing")
}


func main() {
	var p = &Primary{"pAAA", 10, 0.0}
	(*p).Testing()
	(*p).SetScore(90.5)
	(*p).ShowInfo()

	var q = &College{"pBBB", 10, 0.0}
	(*q).Testing()
	(*q).SetScore(90.5)
	(*q).ShowInfo()
}
