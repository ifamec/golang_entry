# Chapter07 Notes - Array & Slice

- Program - maintainable
- Code - readable extendable

### Definition

```
var arrayName [length]dataType
```

`var intArr [3] int`

1. use `&arr` to get the address of the array
2. the address of the first element of the array is the address of the array
3. address of second element is &arr + dataType byte
![array_in_ram](img/array_in_ram.png)
   here, int is 8 bytes

// 4 ways to init an array
```go
var arrInit1 [3] int = [3] int {1, 2, 3}

var arrInit2 = [3] int {4, 5, 6}

var arrInit3 = [...] int {7, 8, 9}

var arrInit4 = [...] int {1: 800, 0: 700, 2: 1000, 4: 1200}

arrInit5 := [...] string {1: "MY", 0: "OH", 2: "GOD"}
```

### Iteration

1. use `for`
2. use for-range
    ```go
    for index, value := range arr {
        BODY
    }
    ```
    - index is the position [scope: for loop]
    - value is the value [scope: for loop]
    - use `_` to ignore index / value
3. 