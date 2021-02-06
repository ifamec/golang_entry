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

	// monster info write to redis
	fmt.Println("Add Goods, Input 'done' to finish")
	for {
		var good string
		fmt.Printf("View Good: ")
		fmt.Scanln(&good)
		if good == "done" {
			fmt.Println("Finish View Goods")
			break
		}

		_, err = conn.Do("LPUSH", "goods", good)
		if err != nil {
			fmt.Println("LPUSH ERROR:", err)
		}
		fmt.Println("Good Recorded")
	}
	fmt.Println()

	recent, err := redis.Strings(conn.Do("LRANGE", "goods", "0", "9"))
	if err != nil {
		fmt.Println("RECENT GOODS ERROR:", err)
	}
	fmt.Println("Recent 10 Goods Viewed:", recent)

}