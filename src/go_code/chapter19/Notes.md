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