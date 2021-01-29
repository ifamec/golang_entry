# Chapter 14 Notes - IO

Handle file as STREAM
- Input Stream (Read)
- Output Stream (Write)

1. Open `os.Open(name string) (file *File, err error)`

## os.File
2. Open `os.File.Close(f *File) error`

e.g.
**Read With Buffer- Big File**
```go
// default buffer is 4096 bytes
reader := bufio.NewReader(file)
for {
    str, err := reader.ReadString('\n') // end by each line
    if err == io.EOF { break }
    fmt.Print(str)
}
```
**Read Whole File - Small File**
- No Open Close
```go
content, err := ioutil.ReadFile(file)
content: []byte
```
**Write File**
`os.OpenFile(name string, flag int, perm FileMode) (file *File, err error)`
flag:
- O_RDONLY: read only
- O_WRONLY: write only
- O_RDWR: read and right
- O_APPEND: append to the end
- O_CREATE: if not exist, create a new file
- O_EXECL: use with O_CREATE, while file not exist
- O_SYNC: open for sync I/O
- O_TRUNC: clean file when open if possible
Use Pipe to combine
  
Write file with buffer `bufio.NewWrite()`

**Read Whole File **
- No Open Close
```go
err := ioutil.ReadFile(file, []byte, perm)
```

**Check if Dir / File exist**
`os.Stat(path String) (fi FileInfo, err error)`
if err == nil - exist
if type of err use os.IsNotExist() == true - not exist
if err type is other - unsure
```go
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}
```

## Copy a file
`io.Copy(dst Writer, src Reader) (written int64, err error)`

## Get Command Args
`os.Args` demo: bin/checkArgs

use flag package demo: bin/checkArgs_flag

## JSON
