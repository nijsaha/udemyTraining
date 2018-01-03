package main

import "fmt"

func main() {

	//when anonymous functions are assigned to a variable it is a func expression
	greeting := func() {
		fmt.Println("hello")
	}
	greeting()
	fmt.Printf("%T\n", greeting)

	greet := makegreet()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

}

func makegreet() func() string {
	return func() string {
		return "returning a function which is returning a string"
	}

}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
