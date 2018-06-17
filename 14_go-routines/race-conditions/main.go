//concurrency is dealing with lots of things at once; whereas paralleism is doing lots of things at once.
//concurrency is independently executing several threads/processes.  doing many things; but only one at a time
//paralleism is simulatenously executing several threads / processes.  doing all at once.

//go run -race main.go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {

	//concurrency - main runs separately in a different thread to foo and bar. without wait, main runs and exists
	wg.Add(2) //adds 2 to the wait group
	go increment("foo: ")
	go increment("bar: ")
	wg.Wait() //waits till the waitgroup runs to 0
	fmt.Println("Final Counter: ", counter)
}

func increment(a string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		//counter++
		fmt.Println(a, counter)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x

	}
	wg.Done() //takes 1 off the waitgroup
}

//race conditions can happen if threads are going for the same resource in parallel/during concurrent access.
//to void we should do
