package main

import "fmt"

func sum(n1 int, args... int) (sum int) {
	sum = n1
	for _, val := range args{
		sum += val
	}
	for i:=0; i<len(args); i++ {
		fmt.Println(args[i])
	}
	return
}

func main()  {
	fmt.Println(sum(1,2,3,4))
}