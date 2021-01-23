# Chapter 05 Notes - Control Flow

Sequence
Choice
Loop


## Choice

### if

```go
if CONDITION {
	
}

if CONDITION {
	
} else {
	
}

if CONDITION_1 {
	
} else if CONDITION_2 {
	
} else {
	
}

```

wrap if CONDITION with `()`, could compile, but not recommend

do not nest too much in choice flow

### switch

NO `BREAK` after each statement

```go
switch CONDITION {
    case CONDITION_1, CONDITION_2:
    	STATEMENT
    case CONDITION_3:
    	STATEMENT
    default:
        STATEMENT
}
```

#### Details
1. CONDITION could be a const variable or a func with return value
2. case CONDITION should have the same data type as switch CONDITION
3. case could have multiple CONDITION
4. if case CONDITION is a const, then cannot dupe
5. no break after each case statement. after the code clock, exit switch after execution. if no case match, go to default
6. default is not necessary
7. switch could have no CONDITION treat as if-else
8. can declare and assign a new variable, NOT RECOMMENDED
9. switch `fallthrough` - continue to next case, NOT RECOMMENDED
10. Type Switch: switch could be used in type-switch to check the actual variable type that interface variable pointing to

## Loop

### for

```go
for (loop variable init); (loop condition); (loop variable iterate) {
	loop body
}
```
1. loop condition is a bool
2. `for condition { body }`, move init and traverse to somewhere else
3. `for {body}`, use with `break` and `continue`
4. for-range: string / array traverse

#### Details
- legacy traverse of string - use byte to iterate 0-255
  - to solve this, use slice
- for-range using character to iterate

### while / do...while

NO WHILE / DO WHILE

use for/if realise

```go
//while:
for {
    if loop condition {
        break
    }
    loop body
    loop iteraton
}
```

do while at least execute once
```go
// do while
for {
    loop body
    loop iteraton
    if loop condition {
        break
    }
}
```

### nested loop

see demo

### break

```go
// to get a random number, set a seed for rand
rand.Seed(time.Now().Unix()) // get a unix time second
// or use UnixNano() 
var num = rand.Intn(100) + 1
```

```go
{ // loop body
    ...
    if CONDITION {
        break
    }
}
```

#### Details
- could use label to indicate jump out of which nested loop
- no label -> end current loop

### continue

- end current loop
- could use label

```go
{ // loop body
  ...
  if CONDITION {
    continue
  }
}
```

## goto / return

### goto

1. shift to certain line
2. use along with condition statement to realise condition shifting, jump out of loop body
3. not recommend to use `goto`

```go
goto lable

lable: CONDITION
```

### return

use in function/method, jump out of current function/method