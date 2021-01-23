package main

import "fmt"

func main()  {
	var score [3][5]float64

	for classNum, class := range score {
		for stuNum, _ := range class {
			fmt.Println("CLASS", classNum + 1, "STUDENT", stuNum + 1)
			fmt.Scanf("%f\n", &score[classNum][stuNum])
		}
	}
	fmt.Println(score)
	sumTotal := 0.0
	stuTotal := 0
	for classNum, class := range score {
		sum := 0.0
		for _, stu := range class {
			sum += stu
			stuTotal ++
		}
		fmt.Println("CLASS", classNum + 1, "SUM IS", sum, "AVG IS", sum / float64(len(class)))
		sumTotal += sum
	}
	fmt.Println("ALL CLASS SUM IS", sumTotal, "AVG IS", sumTotal / float64(stuTotal))

}