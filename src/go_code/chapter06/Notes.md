#Chapter 06 Notes - Function

```go
func funcName (params list) (returnVal dataType list) {
	Code Block
}
// if only have one rtnval, could remove "()"
```

## Package
import functions from other `.go` file.  
package ~~ directory

Golang use package to manage files and project directory
- differentiate same name function
- easy to manage project
- control func variable scope
- `package "PACKAGENAME"` - pack
- `import "PACKAGENAME"` - use
- `PACKAGENAME.FunctionName` 
    - uppercase first character of the fn name
        - which means the fn could be exported
    
### Details

- package name usually == directory name, but not required
- import and use
    - import "pn"
    - import ( "pn1"; "pn2")
- import path start with dirs under `$GOPATH`
- could rename package
    - import alias "packageName" 
- could not dupe function/variable name in the same package (even in different `.go` file)
- main is the package that could be compiled

e.g. ` go build -o bin/function_package go_code/chapter06/function02_package/main
`

the compiled file locates in bin/

## Recursive

a function call itself in its function body

- when execute a function, create a new space in ram
- function scope variables are individual
- recursive need to approaching to exit recursive or, never exit
- when one fn finish, or meet return, function will return to it's caller

### Details

1. params and rtnval list could be multiple
2. dataType of param and rtnval could be varied
3. naming - see variable naming rules
    - start with uppercase - public across packages
    - start with lowercase - private within this packages
4. variables in function are local
5. basic data type and array will pass a copy, will not influence the original value
6. to change the value outside the function, pass address, handle pointer in function
7. NO function override
8. function is a data type, could assign to a variable. we can use that variable to call the function
9. function could be a param
10. go could declare custom data type
    - type custom int // ALIAS
    - e.g. type mySum func(int, int) int
11. could naming return value in func declaration
12. `_` to ignore rtnval
13. go support multiple params
    - `args` is slice use args[index] to visit each val
    - `args...` should be the last param
    
## init

each source file could have one `init()`, which will execute before `main()`

### Details
1. order: global variable definition -> init() -> main()
2. imported package will execute first

## anonymous 

could be single use, or multiple use

- call anonymous function when declare
- assign to variable
- global anonymous function

## closure 

function + it's related env

```go
func accumulator() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}
func main()  {

f := accumulator()
fmt.Println(f(1)) // 11
fmt.Println(f(2)) // 13
fmt.Println(f(3)) // 16
}
```

1. accumulator() will return a func(int) int
2. the anonymous function and `n` became an entity -> closure
3. closure refers as a class, the function is the action, `n` is a field
4. when call `f` function, `n` will init only once, and the `n` is accumulated
5. to scope the closure, need to clarify the variable that the function used

```go
func makeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if ! strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}
```
the closure contains thr `return functon` and `suffix string`

```go
closure := makeSuffix(".jpg")
fmt.Println(closure("aaa"))
fmt.Println(closure("aaa.jpg"))
```
the second time using the closure, we don't need to pass ".jpg" again, since the closure have the suffix value as ".jpg" already

## defer
when creating source, add defer, the source will be released timely

push `defer` statement to a stack, when function finish, pop and execute

when push `defer` statement into the stack, KEEP THE CURRENT VALUE

e.g.:
```psudocode
file = openfile
defer file.close
// --

connect = opendb
defer connect.close
```
the source will be closed before exiting the fn

## Review Params

- pass by value - a copy of the value (basic type, array, struct)
- pass by reference - a copy of the reference (pointer, slice, chan, map, interface)
    - more efficiency copy size smaller
    
pass reference in function to change the value out side the scope

## Scope

- local variable - declare a variable inside a function - scope to current function
- global variable - declare a variable outside a function - scope to current package
    - if a global variable start with uppercase, scope to the whole project
- declare in for/if code block - scope to current code block

assignment cannot execute outside a function

## string
1. `len(str)` - get length [builtin]
2. `[]rune(str)` - iteration string with unicode
3. `n, err := strconv.Atoi("str")` - string to int
4. `n, err := strconv.Itoa("str")` - int to string
5. `[]byte("str")` - string to byte
6. `string([]byte{97,98,99})` - byte to string
7. `strconv.FormatInt(123,2)` - 10 to b2 b8 b16
8. `strings.Contains("string", "str)` - contains substring
9. `strings.Count("cheese","e")//4` - count substring
10. `strings.EqualFold("abc","ABC")` - case insensitive string compare
11. `strings.Index("string", "tri")` - return substring first shows up position
12. `strings.LastIndex("stringstring", "tri")` - return substring last shows up position
13. `strings.Replace("go golang","go", "GO", n)` - replace "oriString" "substring to raplace" "expected value", "how much to replace" **-1, replace all**
14. `strings.Split("Hello,World", ",")` - split string, return an array
15. `strings.ToLower("AAAA") strings.ToUpper("aaaa")` - toLowerCase, toUpperCase
16. `strings.TrimSpace("   Hello, World      ")` - trim spaces
17. `strings.Trim("!!!   Hello, World!! !", " !")` - trim certain charters
18. `strings.TrimLeft("!!!   Hello, World!! !", " !")` - trim certain charters on the left
19. `strings.TrimRight("!!!   Hello, World!! !", " !")` - trim certain charters on the right
20. `strings.HasPrefix("Hello, World", "Hello")` - has prefix
21. `strings.HasSuffix("Hello, World", "World!")` - has suffix


## Time & Date

- use package `time`
- time.Time data type to represent time
- get current time `time.Now()`
- get date info `now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second()`
- format datetime
    - use `Printf` or `Sprintf`
    - use `now.Format("2006-01-02 15:04:05")` could take part of the datetime
        - must use `2006-01-02 15:04:05`
- time const `10*time.Second() // 10s`: // get certain time ***DO NOT USE DIVIDE*
    - Nanosecond
    - Microsecond
    - Millisecond
    - Second
    - Minute
    - Hour
- `time.Sleep()` - sleep
- get Unix (nanosecond) timestamp: `time.Unix() time.Unixnano()`