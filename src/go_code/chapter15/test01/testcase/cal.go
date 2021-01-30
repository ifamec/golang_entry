package main

func getSum(n int) (rtnval int) {
	for i := 1; i <= n; i++ {
		rtnval += i
	}
	return
}

func getDiff(n1 int, n2 int) (rtnval int) {
	return n1 - n2
}