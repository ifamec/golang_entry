package main

import "fmt"

type Stu struct {
	Name string
	Age int
	Address string
}

func main()  {
	var m map[int]int = map[int]int{
		1: 99,
		2: 99,
		3: 99, // need "," here
	}
	fmt.Println(m)
	modify(m)
	fmt.Println(m)

	// use struct as value
	var students map[int]Stu = make(map[int]Stu)
	s1 := Stu{"Tom", 26, "SEA"}
	s2 := Stu{"Jerry", 27, "CAN"}
	students[1] = s1
	students[2] = s2
	fmt.Println(students)

	// traverse
	for k, v := range students {
		fmt.Printf("Num:%v\t Name:%v\t Age:%v\t Address:%v\n", k, v.Name, v.Age, v.Address)
	}

	exercise()
}

func modify(m map[int]int) {
	m[10] = 10
}

func exercise()  {
	var users map[string]map[string]string = map[string]map[string]string{
		"usr1": {"nickname": "NN-usr1"},
	}
	fmt.Println(users)
	modifyUser(users, "usr1")
	fmt.Println(users)
	modifyUser(users, "usr2")
	fmt.Println(users)
	modifyUser(users, "usr3")
	fmt.Println(users)
}
func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = map[string]string{
			"nickname": "nn-" + name,
			"pwd": "888888",
		}
	}
}