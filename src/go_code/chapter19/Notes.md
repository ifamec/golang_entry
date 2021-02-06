# Chapter 19 Notes - Redis

1. NoSQL DataBase 
2. Remote Dictionary Server
3. Data Structure Server

```
Redis Port: 6379

Redis Core:
- resolve command
- handle command

RAM:
- string k-v
- hash
- list
- set
- zset // ordered set

export/import to file `.dbf`, `.dbm`
```

[Redis Doc](https://redisdoc.com)

Redis have 16 default db 0-15

## Intro
1. add k-v `set`
2. check all key `keys *`
3. get value `get key`
4. switch redis db `select index`
5. check current db k-v amount `dbsize`
6. clean current/all db k-v `flushdb | flushall`

redis data type: string, hash, list, set, zset(sorted set)

## string

- key - value
- string type is binary safe, could store img etc
- string value maximum 512MB

```
key: address
value: seattle

```
- `SET address "Seattle WA"`
- `GET address`
- `DEL address`
- setex(set with expire): `SETEX key seconds value`
- mset(set multiple): `MSET key1 val1 keya val2`
- mget(get multiple): `MGET key1 key2`

## Hash
likes map in golang

- key-value pair set
- mapping table of field(string type)
- key is unique

Use: 
- `HSET key field1 value1`  
- `HGET key field`    
- `HGETALL key`  
- `HDEL key field`

Set Get Multiple  
- `HMSET key f1 v1 f2 v2` 
- `HMGET key f1 f2`

- `HLEN key` get length
- `HEXISTS key field` check if field exists

### Practise
```redis
hmset stu1 name tom age 22 score 80 address seattle
```

## List

- ordered by insert position
- could add data in front or rear
- is a liked list, has order, element could be duplicated

Use:
- `lpush key c1 c2 c3` // c3-c2-c1 // push from left
- `lrange key start end(-1 last)`
- `rpush key c1 c2` // push from right
- `lpop key` // pop from left
- `rpop key` //pop from right
- `del key` // delete

Details:
- `lindex key index` get value use index
- `llen key` get list length, if not exist, return 0
- if pop all content, the list gone automatically


## set

string type unordered
- Hashtable data structure
- string element
- unordered
- unique
    - username unique

Use:
- `sadd key val1 val2`
- `smembers key` // get all elements from the set
- `sismember key val` // check if the value is in set or not
- `srem key val` // remove certain value


## redis in golang
lib: https://github.com/gomodule/redigo
e.g. [redisgo](redis01/main/main.go) 

## Redis Pool
- start several connections in redis pool
- pickup and use directly from the pool
- save time

```go
var pool *redis.Pool
pool = &redis.Pool{
	MaxIdle: 8, // maximum connection
	MaxActive: 0, // how many could connect to db, 0 -> unlimited
	IdleTimeout: 100, // maximum timeout
	Dial: func () (redis.Conn, error) {
        return redis.Dial("tcp", "localhost:6379")
	}
}
c := pool.Get()
pool.Close() // once closed, gone forever
```