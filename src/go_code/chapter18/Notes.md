# Chapter 18 Notes - TCP Socket Programming

Network Programming
- TCP Socket Programming
- B/S Http Programming

## Project 01

Server
- listen port 8888
- accept and link client tcp connection
- goroutine to handle request

client
- build connection
- send data(terminal), get value from server
- close connection

=====
server.go
```
listening :8888
    goroutine a - handle A request
    goroutine b - handle B request
    ...
```
====

```
A client.go
request server 8888
```
```
B client.go
request server 8888
```

## Project 02

**Instant Messaging System**

Requirements:
- Register & Login
- Show User List
- Broadcast (Group Chatting)
- P2P Chat
- Offline Message

UI: Terminal

Database: Redis