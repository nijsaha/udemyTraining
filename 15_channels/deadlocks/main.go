package main

import "fmt"

func main() {

	c := make(chan int)
	//c <- 1 //deadlock
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
