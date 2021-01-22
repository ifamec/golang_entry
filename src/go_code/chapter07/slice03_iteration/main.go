package main

import "fmt"

func main()  {
	var arr [5]int = [5]int{1,2,3,4,5}
	var slice []int = arr[1:4]

	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d  ", slice[i])
	}
	fmt.Println()
	for _, value:= range slice {
		fmt.Printf("%d  ", value)
	}
	fmt.Println()

	var slice2 []int = arr[:]
	fmt.Println(slice2)

	var slice3 []int = slice[1:2]
	fmt.Println(slice3)

	slice3[0] = 100
	fmt.Println(arr, "\n",slice, slice2, slice3)


	// append
	fmt.Println()
	fmt.Println()
	fmt.Println()
	var slice4 []int = make([]int, 1, 3)
	slice4[0] = 1
	fmt.Printf("%p\t", &slice4)
	fmt.Println(slice4, &slice4[0], len(slice4), cap(slice4))
	slice4 = append(slice4, 200)
	fmt.Println(); fmt.Println("append:")
	fmt.Printf("%p\t", &slice4)
	fmt.Println(slice4, &slice4[0], &slice4[1], len(slice4), cap(slice4))

	fmt.Println(); fmt.Println("append less than len-cap")
	var slice5 []int = append(slice4, 200)
	fmt.Printf("%p\t", &slice5)
	fmt.Println(slice5, &slice5[0], &slice5[1], &slice5[2], len(slice5), cap(slice5))

	slice4 = append(slice4, 1)
	fmt.Printf("%p\t", &slice4)
	fmt.Println(slice4, &slice4[0], &slice4[1], &slice4[2], len(slice4), cap(slice4), slice5)

	fmt.Println(); fmt.Println("... First time increase capacity: ")
	slice4 = append(slice4, 1)
	fmt.Printf("%p\t", &slice4)
	fmt.Println(slice4, &slice4[0], &slice4[1], len(slice4), cap(slice4))

	slice4 = append(slice4, 200, 300)
	fmt.Printf("%p\t", &slice4)
	fmt.Println(slice4, &slice4[0], &slice4[1])

	fmt.Println(); fmt.Println("... Second time increase capacity: ")
	slice4 = append(slice4, slice4...) // append slice
	fmt.Printf("%p\t", &slice4)
	fmt.Println(slice4, &slice4[0], &slice4[1])

	// copy
	fmt.Println()
	fmt.Println()
	fmt.Println()
	var ori []int = []int{1,2,3,4,5}
	fmt.Println(ori, &ori[0])
	var cop []int = make([]int, 3, 3)
	fmt.Println(cop, &cop[0], len(cop), cap(cop))

	fmt.Println("Destination has less len & cap")
	copy(cop, ori)
	fmt.Println(cop, &cop[0], len(cop), cap(cop))


	fmt.Println("Destination has less len")
	var co  []int = make([]int, 3, 10)
	copy(co, ori)
	fmt.Println(co, &co[0], len(co), cap(co))

	fmt.Println("Destination has more len")
	var c  []int = make([]int, 6, 10)
	copy(c, ori)
	fmt.Println(c, &c[0], len(c), cap(c))
}