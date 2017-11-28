package main

import "fmt"

func main() {
	const (
		A = iota
		B = iota
		C = iota
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	const (
		_  = iota
		KB = 1 << (iota * 10)
		MB = 1 << (iota * 10)
	)
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)

}
