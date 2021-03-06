# Chapter 02 Notes

## Go Execution
Compile then Run: 
- `.go` - go build -> executable file - run -> result 
- could run on no SDK env
- the compiled file contains dependencies, file size would get larger

Run directly (build and compile):
- `.go` - go run -> result
- could not run on no SDK env

### Compile Options

assign build file name: go build -o newfilename x.go

## ! Hints
1. go file end with `.go`
2. execution entry: `main()`
3. case-sensitive
4. no `;` at the end of each line
5. cannot combine multiple statements in one line
6. compile won't pass if imported package / declared variable is not used
7. {} need to show up in pair

## Escape Char

\r will replace the front string with back string byte by byte

## Comment

- `//   `   - comment a line
- `/* */`   - comment a code block

## Coding Style
- use line comment in developing as much as possible  
- `gofmt -w x.go` could format the code
- Try limit 80 characters in one line

## 