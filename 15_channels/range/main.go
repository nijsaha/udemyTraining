package main

import "fmt"

func main() {
	c := make(chan int) //unbuffered channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i //gets blocked till the channel is cleared.  Heere in this e.g range takes it and clears the channel
		}
		close(c) //channel closes only after the completion of for loop
	}()

	for n := range c { //halts over here until the channel is closed
		fmt.Println(n)
	}
}
