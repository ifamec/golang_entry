# Chapter 16 Notes

## goroutine

e.g. 1 - 20000 find prime
1. for loop
2. use concurrency -> distribute to multiple goroutines
```go
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i * i <= n; i ++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}
```

### Process & Thread
- Process: the execution of a program
- Thread: the unit of execution within a process 
- one process can have multiple threads
- one program at least have one process, one process at least have one thread

### concurrence & parallel
concurrence 
- multiple thread runs one single core
- at certain moment, only one thread is running

parallel 
- multiple thread runs on multiple cores
- at certain moment, multiple is running 

### Goroutine
goroutine - (lightweight thread)  
main thread - more like a process

#### Feature
1. have individual stack space
2. share program heap space
3. user could dispatch
4. goroutine is lightweight thread

if main thread ends while a goroutine not end, quit.
- the main thread is a physical thread, heavy
- goroutine was started from main thread, lightweight
- go could start thousands of goroutine

#### MPG model
- M: main thread
- P: env for goroutine (source etc.)
- G: goroutine

TODO NYI CHECK MPG MODEL EXPLANATION

#### Setup CPU cores
`runtime`
- `runtime.NumCPU()` // check cpu cores
- `runtime.GOMAXPROCS(num)` // setup cpu cores

## channel

e.g. calc ! 1-200
concurrence & parallel safety

> Check if have source race in the program: -race

multiple goroutines write to one Map, causing concurrent error
- use scope variable and lock
    - let each goroutine queue then write to the map
    - `sync.lock`
    - [Check Code](goroutine03/main/main.go)
- use channel

**WHY CHANNEL**
- Cannot estimate sleep time
- sleep time may either too long or too short
- not good for multi-goroutine RW a global variable
- USE CHANNEL !!!

### Intro
- Channel is a Queue
    - first in first out
    - `push()` `pop()`
- Data : first in first out (FIFO)
- thread safe
- channel has type, string channel could only save string type

### Use
- `var channelName chan dataType`
- channel is reference type
- need initiation | use after make
- channel has type
- Write `channel <- data`
- Read `<- channel`  
[Check Code](channel01/main/main.go)
  
### Hints
- Channel could only store certain type of data
- cannot add data if channel is full
    - could add if some data was taken out
- cannot get data if channel is empty // deadlock error

If want to combine different data type data into one channel, declare a `interface{}` channel.   
To get the data (like a struct variable), use type assertion to get the real object.

### Close and Traverse

#### Close
when channel get closed, cannot write into the channel, but could read from channel
[Check Code](channel04_close/main/main.go)

#### Traverse
1. if channel is not closed, will report deadlock error
2. if channel is closed, everything will be good

### Block
comment go reaData() in [Code](channel06_exercise/main/main.go):  
only write no read, causing deadlock  
if only one goroutine is reading, [OK] frequency does not matter

### Details

1. Channel can be read only
    - write only `var chan chan <- int`
    - read only `var chan <- chan int`
    - e.g. define a two-way channel, when pass into a function, defined W/R only in function parameters
    - [Code](channel09/detail01/main.go)
2. `select` solve block issue
    - [Code](channel09/detail02/main.go)
3. use `recover` to capture `panic`
    - [Code](channel09/detail03/main.go)
