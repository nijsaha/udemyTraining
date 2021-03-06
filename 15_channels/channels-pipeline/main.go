package main

import "fmt"

func main() {
	//set up the pipeline
	c := gen(2, 3)
	out := sq(c)

	//consume the output
	fmt.Println(<-out)
	fmt.Println(<-out)

}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
