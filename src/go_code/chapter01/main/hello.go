// the go file belongs to main package, each go file should belongs to a package
package main
// import a package
import "fmt"
// func to indicate main is a function
func main() { // main is the program entry
	fmt.Println("Hello, World!") // call Println method, log the string
}
// compile and run -> go build hello.go
// run directly -> go run hello.go