package main

import "fmt"

func main()  {
	label2:
	for i := 0; i < 4; i++ {
		//label1: // set a label
		for j := 0; j < 10; j++ {
			if j == 2 {
				//continue // jump this loop body // 013456789 *4
				//continue label1 // 013456789 *4
				continue label2 // 01010101
			}
			fmt.Printf("%d", j)
		}
	}
	fmt.Printf("\n\n")

	fmt.Println("EXERCISE")
	exercise()

}
func exercise()  {
	// 1 - 100 odd number
	for i := 1; i <=100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// get input int count even/odd end if 0
	var pos, neg int
	var num int
	for {
		fmt.Println("INPUT:")
		fmt.Scanf("%d\n", &num)
		if num == 0 {
			break // end loop
		}
		if num > 0 {
			pos++
			continue // end current cycle
		}
		neg++
	}
	fmt.Println("pos:", pos, "\tneg:", neg)

	// 100k if remain>50k pay 5% else pay 1000
	var remain float32 = 100000
	var time int
	for {
		if remain < 1000 {
			break
		}
		if remain > 50000{
			remain *= 0.95
			time++
		} else {
			remain -= 1000
			time++
		}
	}
	fmt.Println(time)
}