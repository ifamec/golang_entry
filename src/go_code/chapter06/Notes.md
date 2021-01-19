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