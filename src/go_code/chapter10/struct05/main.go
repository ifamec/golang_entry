package main

import "fmt"

type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp, rightDown Point
}
type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1 has 4 int continuously allocated in ram
	fmt.Printf("r1.leftUp.x: %p\nr1.leftUp.y: %p\nr1.rightDown.x: %p\nr1.rightDown.y: %p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	fmt.Println()
	// r2 has two pointer type, the address of these two pointer is identical
	// the two pointer might not continuous
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp: %p\nr2.rightDown: %p\n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.leftUp Pointing: %p\nr2.rightDown Pointing: %p\n", r2.leftUp, r2.rightDown)
}
