// you have to import "fmt" package
// Package fmt implements formatted I/O with functions analogous to C's printf and scanf. 
// https://pkg.go.dev/fmt

package main
import "fmt"

func main(){

var firstName string
var lastName string
 
  fmt.println("What's your first name?\n")
  fmt.scan(&firsName)
  
  fmt.println("What's your last name?\n")
  fmt.scan(&lastName)
  
  fmt.printf("Hi %v %v, how are you today?", firstName, lastName)

}
