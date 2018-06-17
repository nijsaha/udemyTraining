//using semaphore.
//Semaphone is some kind of flag used to control system resources defined by the programmer and used based on certain conditions

package main

import (
	"fmt"
)

//many functions writing to 1 channel

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
