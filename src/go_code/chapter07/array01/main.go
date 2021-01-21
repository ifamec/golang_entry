package main

import "fmt"

func main()  {
	formatFloat := fmt.Sprintf("%.2f", 3.1415)
	fmt.Println(formatFloat)

	// Declare
	var arr [6]float64
	// Assign Value
	arr[0] = 1.5
	arr[1] = 2.0
	arr[2] = 0.9
	arr[3] = 3.2
	arr[4] = 2.3
	arr[5] = 1.8
	var sum float64
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("%.2f, %.2f\n", sum, sum/float64(len(arr)))

	var intArr [3] int
	fmt.Println(intArr)
	fmt.Printf("%p\n", &intArr)
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// 4 ways to init an array
	var arrInit1 [3] int = [3] int {1, 2, 3}
	fmt.Println(arrInit1)

	var arrInit2 = [3] int {4, 5, 6}
	fmt.Println(arrInit2)

	var arrInit3 = [...] int {7, 8, 9}
	fmt.Println(arrInit3)

	var arrInit4 = [...] int {1: 800, 0: 700, 2: 1000, 4: 1200}
	fmt.Println(arrInit4)

	arrInit5 := [...] string {1: "MY", 0: "OH", 2: "GOD"}
	fmt.Println(arrInit5)

	//arrayExc()
}

func arrayExc()  {
	var arr [5] float64
	for i := 0; i<5; i++ {
		fmt.Printf("INPUT:")
		fmt.Scanf("%f\n", &arr[i])
	}
	fmt.Println(arr)
}