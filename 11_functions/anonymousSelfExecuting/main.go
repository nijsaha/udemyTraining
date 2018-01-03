package main

import "fmt"

func main() {
	//not recommended. Not idiomatic way to write functions in golang
	func() { //anonymous function - no parameters, no identifier, no return type
		fmt.Println("I am self driving")
	}() //this self executes
}
