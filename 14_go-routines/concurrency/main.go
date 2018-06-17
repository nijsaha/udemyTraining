//concurrency is dealing with lots of things at once; whereas paralleism is doing lots of things at once.
//concurrency is independently executing several threads/processes.  doing many things; but only one at a time
//paralleism is simulatenously executing several threads / processes.  doing all at once.

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//no concurrency
	//foo()
	//bar()

	//concurrency - main runs separately in a different thread to foo and bar. without wait, main runs and exists
	wg.Add(2) //adds 2 to the wait group
	go foo()
	go bar()
	wg.Wait() //waits till the waitgroup runs to 0
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done() //takes 1 off the waitgroup
}
func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done() //takes another 1 off the WaitGroup
}
