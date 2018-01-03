package main

import "fmt"

func world() {
	fmt.Println("world")

}
func hello() {
	fmt.Print("hello")

}

func main() {
	defer world() //defers the execution till the last moment when it is about to exit the function it is called in (i.e. main in this instance)
	hello()

	//good use - when a file is openend, remember to close it.
	//insread now defer the closing immediately after opening so that it does close before leaving the function

}
