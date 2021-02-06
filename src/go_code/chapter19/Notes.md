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
- `HSET value field1 value1`  
- `HGET value field`    
- `HGETALL value`  
- `HDEL value field`

Set Get Multiple  
- `HMSET value f1 v1 f2 v2` 
- `HMGET value f1 f2`

- `HLEN value` get length
- `HEXISTS value field` check if field exists

### Practise
```redis
hmset stu1 name tom age 22 score 80 address seattle
```

