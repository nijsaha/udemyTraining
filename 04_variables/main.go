package main

import "fmt"

var abc string

func main() {
	a := 10
	b := "golang"
	c := 4.12
	d := true
	var e = 5

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)
	fmt.Printf("%v - %T \n", e, e)
	hellofunc()
	fmt.Println(pqr)
	incremental := Wrapper()
	fmt.Println(incremental())
}

var pqr float64

func hellofunc() {
	fmt.Println(abc)
	fmt.Println(pqr)
	x := 0

	//anonymous function. A function inside a function. Basically function without a name
	increment := func() int {
		x++
		return x
	}
	fmt.Printf("This is increment of x %d \n", increment())
}

//Wrapper function return types
func Wrapper() func() int {
	x := 20
	return func() int {
		x++
		return x
	}
}
