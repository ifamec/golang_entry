package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "string广州😊"
	fmt.Println(len(str)) // 6 + 3 + 3 + 4 (Chinese Characters - 3 bytes | emoji - 4 bytes)

	for _, v := range []rune(str) {
		fmt.Printf("%c ", v)
	}
	fmt.Println("")

	n1, err := strconv.Atoi("33")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n1)
	}

	n2 := strconv.Itoa(12345)
	fmt.Printf("%T - %v\n", n2, n2)

	fmt.Println([]byte("Hello"))

	fmt.Println(string([]byte{97, 98, 99}))

	fmt.Println(strconv.FormatInt(28, 2), strconv.FormatInt(28, 8), strconv.FormatInt(28, 16))

	fmt.Println(strings.Contains("string", "tri"), strings.Contains("string", "tir"))

	fmt.Println(strings.Count("cheese", "e"))

	fmt.Println(strings.EqualFold("abc", "ABC"), "abc" == "ABC")

	fmt.Println(strings.Index(str, "州"), strings.Index(str, "Hello"))
	fmt.Println(strings.Index("stringstring", "tri"))

	fmt.Println(strings.LastIndex("stringstring", "tri"))

	fmt.Println(strings.Replace("go go golang","go", "GO", -1))

	fmt.Println(strings.Split("Hello,World,OK", ","))

	fmt.Println(strings.ToLower("AAAA"), strings.ToUpper("aaaa"))

	fmt.Println(strings.TrimSpace("   Hello, World      "))

	fmt.Println(strings.Trim("!!!   Hello! World!! !", " !"))

	fmt.Println(strings.TrimLeft("!!!   Hello, World!! !", " !"))

	fmt.Println(strings.TrimRight("!!!   Hello, World!! !", " !"))

	fmt.Println(strings.HasPrefix("Hello, World", "Hello"))

	fmt.Println(strings.HasSuffix("Hello, World", "World!"))


}
