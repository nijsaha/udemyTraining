package main

import "fmt"

func main() {
	if true {
		fmt.Println("always true")
	}
	if false {
		fmt.Println("this did not run")
	}

	b := true
	if food := "chocolate"; b {
		fmt.Println(food)
	}

	if !true {
		fmt.Println("first statement")

	} else {
		fmt.Println("second statement")
	}

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			println("hello i (", i, ") is divisible by 3")
		}
	}
}
