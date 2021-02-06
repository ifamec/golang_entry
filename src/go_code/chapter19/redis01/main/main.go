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
	_, err = conn.Do("SET", "name", "Tom")
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	fmt.Println("Set Success")

	// Read Data
	val, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	// val is interface{}, since name is a string
	fmt.Println("Get Success, Value is", val)

}