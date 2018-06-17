/*
Running factorial 100 times in parallel and concurrently
*/
package main

import "fmt"

func main() {

	in := numgen()
	f := factorial(in)
	for n := range f {
		fmt.Println(n)
	}

}
func numgen() chan int {
	out := make(chan int)
	go func() {
		for i := 4; i <= 14; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

//sending multiple factorial
func factorial(in chan int) chan int64 {
	out := make(chan int64)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

//Actual factorial
func fact(num int) int64 {
	var total int64
	total = 1
	for i := num; i > 0; i-- {
		total *= int64(i)
	}
	return total
}
