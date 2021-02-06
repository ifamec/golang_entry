package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

// init pool when execute
func init()  {
	pool = &redis.Pool{
		MaxIdle: 8, // maximum connection
		MaxActive: 0, // how many could connect to db, 0 -> unlimited
		IdleTimeout: 100, // maximum timeout
		Dial: func () (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("SET", "namepool", "RRR")
	if err != nil {
		fmt.Println(err)
	}
	val, _ := redis.String(c.Do("GET", "namepool"))
	fmt.Println(val)

	pool.Close()
	// c2 := pool.Get()
	// fmt.Println(c2)
	// _, err = c2.Do("SET", "namepool2", "RRR22")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// val2, _ := redis.String(c2.Do("GET", "namepool2"))
	// fmt.Println(val2)
}
