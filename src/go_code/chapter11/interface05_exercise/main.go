package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Realise Sort for slice of struct "Hero" sort.Sort(data Interface)
// declare a struct
type Hero struct {
	Name string
	Age int
}
// slice of Hero
type HeroSlice []Hero

// realise Interface interface
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool { // Sort by what small->large or ..
	// age small to large
	return hs[i].Age < hs[j].Age // sort by age
	// return hs[i].Name < hs[j].Name // sort by name
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

// STUDENT SLICE
type Student struct {
	Name string
	Age int
	Score float64
}// realise Interface interface
type StudentSlice []Student
func (s StudentSlice) Len() int {
	return len(s)
}
func (s StudentSlice) Less(i, j int) bool {
	return s[i].Score < s[j].Score // sort by age
}
func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}



func main() {

	// Sort Slice
	var intSlice []int = []int{10, -1, 8, 29, -32}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// sort Hero Slice
	var hs HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hn_%d", rand.Intn(100)),
			Age: rand.Intn(1000),
		}
		hs = append(hs, hero)
	}
	for _, v := range hs {
		fmt.Printf("%v\t", v)
	}
	sort.Sort(hs)
	fmt.Println()
	for _, v := range hs {
		fmt.Printf("%v\t", v)
	}

	fmt.Println()
	fmt.Println()
	var stu StudentSlice
	for i := 0; i < 10; i++ {
		s := Student{
			Name: fmt.Sprintf("stu_%d", rand.Intn(100)),
			Age: rand.Intn(30),
			Score: rand.Float64() * 100,
		}
		stu = append(stu, s)
	}
	for _, v := range stu {
		fmt.Printf("%v\n", v)
	}
	sort.Sort(stu)
	fmt.Println()
	for _, v := range stu {
		fmt.Printf("%v\n", v)
	}
}
