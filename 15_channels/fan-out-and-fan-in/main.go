package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	//FAN out
	//Distribute the sq work  across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	//FAN in
	//Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n) //4 then 9; or 9 then 4
	}

}
func gen(nums ...int) chan int {
	out := make(chan int)
	fmt.Printf("Type of nums: %T\n", nums)
	go func() {
		//for key,value:=range slice  - Note range over slice brings in key value
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
		for n := range in { //range over channel just brings value
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	fmt.Printf("Type of cs: %T\n", cs)

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs { //slice of chan int
		go func(ch chan int) { //this is the function parameters
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c) //note this is the function argument being passed in. i.e function call
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
