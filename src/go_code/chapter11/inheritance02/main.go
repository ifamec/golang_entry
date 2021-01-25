package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

func (stu *Student) ShowInfo() {
	fmt.Println("Name:", (*stu).Name, "Age:", (*stu).Age, "Score", (*stu).Score)
}
func (stu *Student) SetScore(s float64) {
	(*stu).Score = s
}
func (stu *Student) GetSum(m int, n int) int {
	return n + m
}

// Primary
type Primary struct {
	Student // embed Student anonymous fields
}

func (p *Primary) Testing() {
	fmt.Println("Primary Testing...")
}

// College
type College struct {
	Student
}

func (p *College) Testing() {
	fmt.Println("College Testing...")
}

func main() {
	var p = &Primary{}
	(*p).Student.Name = "pAAA"
	(*p).Student.Age = 10
	(*p).Testing()
	(*p).Student.SetScore(90.5)
	fmt.Println((*p).Student.GetSum(1, 2))
	(*p).Student.ShowInfo()

	var c = &College{}

	(*c).Testing()
	(*c).SetScore(98.5)
	fmt.Println((*c).Student.GetSum(1, 2))
	(*c).ShowInfo()
}
