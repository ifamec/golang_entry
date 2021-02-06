package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

func main()  {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Redis Conn Error:", err)
	}
	defer conn.Close()

	// monster info write to redis
	var i int = 0
	for {
		if i > 2 {
			break
		}
		// get user input
		i ++
		var name, age, skill string
		fmt.Printf("Add Monster: ")
		fmt.Scanln(&name, &age, &skill)

		_, err = conn.Do("HMSET", "monster_" + strconv.Itoa(i), "name", name, "age", age, "skill", skill)
		if err != nil {
			fmt.Println("HMSET ERROR:", err)
		}
		fmt.Println("Monster Add Success")
	}
	fmt.Println("Monster Add Finish")
	fmt.Println()

	keys, err := redis.Strings(conn.Do("keys", "monster*"))
	if err != nil {
		fmt.Println("Get Keys Error", err)
	}
	fmt.Println("Get Keys Success:", keys)
	for _, v := range keys{
		all, err := redis.Strings(conn.Do("HMGET", v, "name", "age", "skill"))
		if err != nil {
			fmt.Println("HMGET ERROR:", err)
		}
		fmt.Println(all)
	}

}