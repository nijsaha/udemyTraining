package main

import "fmt"

const s = "Hello World"
const (
	hello = "Hello"
	base  = "0"
)

func main() {
	const pi = 4.12
	fmt.Printf("%f\n", pi)
	fmt.Println(base)
}
