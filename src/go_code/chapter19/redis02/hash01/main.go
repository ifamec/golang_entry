package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Redis Conn Error:", err)
	}
	defer conn.Close()

	// Write Data
	_, err = conn.Do("HSET", "u01", "name", "Tom")
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	_, err = conn.Do("HSET", "u01", "age", "16")
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	fmt.Println("Set Success")

	// Read Data
	name, err := redis.String(conn.Do("HGET", "u01", "name"))
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	age, err := redis.Int(conn.Do("HGET", "u01", "age"))
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	// val is interface{}, since name is a string
	fmt.Println("Get Success : Name is", name + ",", "Age is", age)

}