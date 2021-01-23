package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//ex1()
	//ex2()
	//ex3()
	//ex4()
	//ex5()
	//ex6()
	//ex7()
	//ex8()
	//ex9()
	ex10()
}

func ex1() {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	for i, _ := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("ORI", arr)
	for i, _ := range arr {
		if i > len(arr)/2 {
			break
		}
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("REVERSE", arr)
	sum, max, maxIdx, target := 0, -1, -1, 55
	isContain := false
	for index, val := range arr {
		sum += val
		if val > max {
			max = val
			maxIdx = index
		}
		if val == target {
			isContain = true
		}
	}
	fmt.Println("AVG", float64(sum)/float64(len(arr)),
		"    MAX IS", max,
		"    INDEX OF MAX IS", maxIdx,
		"    IS CONTAIN 55", isContain)

}
func ex2() {
	var arr [5]int = [5]int{1, 3, 5, 7, 9}
	var insertNum int = 1
	var rst [6]int

	// do insertion
	var slice []int = make([]int, 0, 6)
	if insertNum <= arr[0] {
		slice = append(slice, insertNum)
		slice = append(slice, arr[:]...)
	} else if insertNum >= arr[len(arr)-1] {
		slice = append(slice, arr[:]...)
		slice = append(slice, insertNum)
	} else {
		var idx int = 0
		// find insertion point
		for i := 0; i < len(arr); i++ {
			if insertNum <= arr[i] {
				idx = i
				break
			}
		}
		slice = append(slice, arr[:idx]...)
		slice = append(slice, insertNum)
		slice = append(slice, arr[idx:]...)
	}
	//fmt.Println(slice)
	for i, v := range slice {
		rst[i] = v
	}
	fmt.Println(rst)

}
func ex3() {
	var arr [3][4]int
	for rowIndex, row := range arr {
		for colIndex, _ := range row {
			fmt.Println("INPUT ROW", rowIndex, "COL", colIndex)
			fmt.Scanf("%d\n", &arr[rowIndex][colIndex])
		}
	}
	fmt.Println(arr)
	for rowIndex, row := range arr {
		for colIndex, _ := range row {
			if rowIndex == 0 || colIndex == 0 || rowIndex == len(arr)-1 || colIndex == len(row)-1 {
				arr[rowIndex][colIndex] = 0
			}
		}
	}
	fmt.Println(arr)
}
func ex4() {
	var arr [4][4]int
	for rowIndex, row := range arr {
		for colIndex, _ := range row {
			fmt.Println("INPUT ROW", rowIndex, "COL", colIndex)
			fmt.Scanf("%d\n", &arr[rowIndex][colIndex])
		}
	}
	fmt.Println(arr)
	arr[0], arr[3] = arr[3], arr[0]
	arr[1], arr[2] = arr[2], arr[1]
	fmt.Println(arr)
}
func ex5() {
	var arr [5]int = [5]int{1, 3, 5, 7, 9}
	fmt.Println("ORI", arr)
	for i, _ := range arr {
		if i > len(arr)/2 {
			break
		}
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("REVERSE", arr)
}
func ex6() {
	var arr [10]string = [10]string{"AA", "BB", "CC", "AA", "BB", "CC", "AA", "BB", "CC", "AA"}
	var index []string = make([]string, 0, 2)
	var target string = "AA"

	for i, v := range arr {
		if v == target {
			index = append(index, strconv.Itoa(i))
		}
	}
	fmt.Println("INDEX OF ALL 'AA':", index)
}
func ex7() {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	for index, _ := range arr {
		arr[index] = rand.Intn(100)
	}
	fmt.Println("ORI", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("SORTED", arr)
	var target = 90
	binarySearch(&arr, 0, len(arr)-1, target)

}
func binarySearch(arr *[10]int, left int, right int, target int) {
	if left > right {
		fmt.Println("NOT FOUND")
		return
	}
	middle := (left + right) / 2
	if (*arr)[middle] > target {
		binarySearch(arr, 0, middle-1, target)
	} else if (*arr)[middle] < target {
		binarySearch(arr, middle+1, right, target)
	} else {
		fmt.Println("FOUND", target, "INDEX IS", middle)
	}
}
func ex8() {
	var arr [5]int
	for i, _ := range arr {
		fmt.Println("INPUT INT:")
		fmt.Scanf("%d\n", &arr[i])
	}
	fmt.Println(arr)
	max, min := arr[0], arr[0]
	maxIndex, minIndex := 0, 0
	for index, value := range arr {
		if value > max {
			max = value
			maxIndex = index
		}
		if value < min {
			min = value
			minIndex = index
		}
	}
	fmt.Println("MAX IS", max, "MAX INDEX IS", maxIndex)
	fmt.Println("MIN IS", min, "MIN INDEX IS", minIndex)
}
func ex9() {
	//var arr [8]int = [8]int{60, 53, 61, 80, 35, 74, 30, 56}
	var arr [8]int = [8]int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(arr)
	sum := 0
	greaterThanAveCount, smallerThanAveCount := 0, 0
	for _, val := range arr {
		sum += val
	}
	var avg float64 = float64(sum) / float64(len(arr))
	fmt.Println("SUM", sum, "AVG", avg)
	for _, val := range arr {
		if float64(val) > avg {
			greaterThanAveCount++
		} else if float64(val) < avg {
			smallerThanAveCount++
		}
	}
	fmt.Println("THE ARRAY HAS", greaterThanAveCount, "NUMBER THAT GREATER THAN AVG")
	fmt.Println("THE ARRAY HAS", smallerThanAveCount, "NUMBER THAT SMALLER THAN AVG")
}
func ex10() {
	var arr [8]float64 = [8]float64{8.5, 1.5, 9.4, 3.3, 4.5, 6.7, 10.0, 8.8}
	//var arr [8]float64 = [8]float64{0, 0, 0, 0, 0, 10, 10, 10}
	max, min := arr[0], arr[0]
	maxIndex, minIndex := 0, 0
	for index, value := range arr {
		if value > max {
			max = value
			maxIndex = index
		}
		if value < min {
			min = value
			minIndex = index
		}
	}
	fmt.Println("HIGHEST IS", maxIndex+1, "LOWEST IS", minIndex+1)
	sum := .0
	for index, value := range arr {
		if index != maxIndex && index != minIndex {
			sum += value
		}
	}
	fmt.Println("SUM", sum)
	avg := sum / (float64(len(arr)) - 2.0)
	fmt.Println("FINAL SCORE IS", avg)
	bestDelta, worstDelta := math.Abs(arr[0]-avg), math.Abs(arr[0]-avg)
	bestIndex, worstIndex := 0, 0
	for index, value := range arr {
		delta := math.Abs(value - avg)
		if delta < bestDelta {
			bestDelta = delta
			bestIndex = index
		}
		if delta > worstDelta {
			worstDelta = delta
			worstIndex = index
		}
	}
	fmt.Println("BEST IS", bestIndex+1, "DELTA IS ONLY", bestDelta)
	fmt.Println("WORST IS", worstIndex+1, "DELTA IS", worstDelta)

	fmt.Println("EXCEPT INVALID SCORE JUDGE:")
	bestDelta, worstDelta = math.Abs(arr[0]-avg), math.Abs(arr[0]-avg)
	bestIndex, worstIndex = 0, 0
	for index, value := range arr {
		if index != maxIndex && index != minIndex {
			delta := math.Abs(value - avg)
			if delta < bestDelta {
				bestDelta = delta
				bestIndex = index
			}
			if delta > worstDelta {
				worstDelta = delta
				worstIndex = index
			}
		}
	}
	fmt.Println("BEST IS", bestIndex+1, "DELTA IS ONLY", bestDelta)
	fmt.Println("WORST IS", worstIndex+1, "DELTA IS", worstDelta)
}
