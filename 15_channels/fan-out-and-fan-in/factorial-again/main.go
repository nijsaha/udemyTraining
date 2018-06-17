/*
Running factorial 1000 times in parallel and concurrently
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	in := numgen()

	//fanning out - using multiple functions to process one channel
	f0 := factorial(in)
	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)
	f6 := factorial(in)
	f7 := factorial(in)

	//fanning in - merging onto 1 channel
	//multiplexing multiple channels onto a single channel
	for n := range merge(f0, f1, f2, f3, f4, f5, f6, f7) {
		fmt.Println(n)
	}

}
func numgen() chan int {
	out := make(chan int)
	go func() {
		for j := 0; j < 10000; j++ { //10,000
			for i := 3; i <= 14; i++ { //10 - total 100,000
				out <- i
			}
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

func merge(cs ...chan int64) chan int64 {
	var wg sync.WaitGroup
	out := make(chan int64)

	//func expression needs definition first before being called in the loop down
	output := func(c <-chan int64) {
		for n := range c { //merging is happening here.
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs { //loopint through the list of channels and merging them
		go output(c)
	}

	//start a goroutine to close out once all the output goroutines are done.
	//this must start after the wg.Add call
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
