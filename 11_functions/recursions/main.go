package main

import "fmt"

func main() {
	//4*3*2*1
	fmt.Println(factorial(4))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
