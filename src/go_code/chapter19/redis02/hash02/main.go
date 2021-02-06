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
	_, err = conn.Do("HMSET", "u02", "name", "Tom", "age", "16", "address", "seattle")
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	fmt.Println("Set Success")

	// Read Data
	values, err := redis.Strings(conn.Do("HMGET", "u02", "name", "age", "address"))
	if err != nil {
		fmt.Println("DO ERROR:", err)
	}
	// val is interface{}, since name is a string
	fmt.Println("Get Success : Vals")
	for v, i := range values {
		fmt.Println(i, v)
	}

}