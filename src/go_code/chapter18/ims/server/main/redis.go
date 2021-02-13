package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

// init pool when execute
func initPool(address string, maxIdle, maxActive int, idle time.Duration)  {
	pool = &redis.Pool{
		MaxIdle: maxIdle, // maximum connection
		MaxActive: maxActive, // how many could connect to db, 0 -> unlimited
		IdleTimeout: idle, // maximum timeout
		Dial: func () (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}