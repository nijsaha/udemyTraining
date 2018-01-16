package main

import (
	"fmt"
)

func main() {
	foo(1, 2, 3)
	foo(1, 2)
	aslice := []int{1, 2, 3, 4, 5}
	foo(aslice...)

}

func foo(numbers ...int) int {
	fmt.Print(numbers)
	return 0
}
