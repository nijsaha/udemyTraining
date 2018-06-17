package main

import "fmt"

func main() {
	c := factorial(14)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(num int) chan int {
	out := make(chan int)
	total := 1
	go func() {
		for i := num; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
