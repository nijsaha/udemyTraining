//Mutex - Mutually Exclusive.
//No matter how many threads are like trying to share a resource,
//the resource can be locked when used by the threads

//atomic - similar to mutex - but very granular.. i.e specifc to a particular object.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64 //note. changed to int64 for doing atomic functions.

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
		atomic.AddInt64(&counter, 1)
		fmt.Println(a, i, "Counter: ", counter)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	}
	wg.Done() //takes 1 off the waitgroup
}

//race conditions can happen if threads are going for the same resource in parallel/during concurrent access.
//to void we should do
