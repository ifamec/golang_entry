# Chapter 15 Notes - Unit Test

Legacy Way cons:
- need main function modification, will infect the running program
- difficult to manage all tests

**Unit Test**

testing framework `go test`  
realise unit test and efficiency
- make sure each function is executable and result is correct
- code efficiency is good
- could find logic error - easy to locate issue, find design issue to maintain stable under high concurrency

## Details

- test file naming
    - original.go
    - original_test.go
- include `func Test/^[A-Z]\w.*/() {}` test function
- *testing.T functions
    - Fatalf
    - Logf
- Could contain multiple test function in one file
- execute
    - `go test`    - print log only if test fails
    - `go test -v` - print log all the time
- use `Fatalf` to print test fail message and exit 
- `Logf` print logs
- testing framework handle all tests outside of main() function
- test single file
    - `go test -v cal_test.go cal.go`
- test single function
    - `go test -v -test.run TestGetSum`
