package main

import "fmt"

func main() {
	var day byte
	fmt.Printf("Input(a-c, 1byte):")
	fmt.Scanf("%c\n", &day)

	switch day {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	default:
		fmt.Println("Error")
	}

	var n1 int32 = 20
	var n2 int64 = 20
	var n3 int32 = 31
	switch n1 {
	//case n2:
	//	fmt.Println(n2) // error
	case 20, 30 + 1:
		fmt.Println("ok", n2)
	//case 31: // DUPE with 30 + 1 above
	//	fmt.Println("ok", n2)
	case n3: // DUPE, but if it is a variable, then could pass compile
		fmt.Println("ok~")
	}

	var age int = 10
	switch {
	case age == 10:
		fmt.Println("10")
	case age == 20:
		fmt.Println("20")
	default:
		fmt.Println("default")
	}

	var score int = 10
	switch {
	case score > 90:
		fmt.Println("A")
	case score >60 && score <= 90:
		fmt.Println("P")
	default:
		fmt.Println("F")
	}

	switch grade := 91; {
	case grade > 90:
		fmt.Println("A~")
	case grade >60 && grade <= 90:
		fmt.Println("P~")
	default:
		fmt.Println("F~")
	}

	var num int = 10
	switch num {
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
		fallthrough
	case 30:
		fmt.Println("30")
	default:
		fmt.Println("40")
	}
	exercise()
}

func exercise()  {
	// get lowercase print uppercase limit to abcde
	var a byte
	fmt.Println("Input(a-e, 1byte):")
	fmt.Scanf("%c\n", &a)
	switch a {
	case 'a': fmt.Println("A")
	case 'b': fmt.Println("B")
	case 'c': fmt.Println("C")
	case 'd': fmt.Println("D")
	case 'e': fmt.Println("E")
	default:  fmt.Println("other")
	}

	// get score [0,100] pass if > 60 else, fail
	var b float64
	fmt.Printf("Input([0,100], 1byte):")
	fmt.Scanf("%f\n", &b)
	switch b>=0&&b<=100 {
	case b>=60: fmt.Println("PASS")
	default:	fmt.Println("FAIL")
	}
	// get month, print season
	var c byte
	fmt.Printf("Input([1,12], 1byte):")
	fmt.Scanf("%d\n", &c)
	switch c {
	case 3, 4,5: fmt.Println("Spring")
	case 6, 7,8: fmt.Println("Summer")
	case 9, 10,11: fmt.Println("Autumn")
	case 12, 1,2: fmt.Println("Winter")
	default: fmt.Println("Error")
	}
	// get string print related
	var d string
	fmt.Printf("Input(string aa-gg):")
	fmt.Scanf("%s\n", &d)
	switch d {
	case "aa": fmt.Println("dish: aa")
	case "bb": fmt.Println("dish: bb")
	case "cc": fmt.Println("dish: cc")
	case "dd": fmt.Println("dish: dd")
	case "ee": fmt.Println("dish: ee")
	case "ff": fmt.Println("dish: ff")
	case "gg": fmt.Println("dish: gg")
	}
}
