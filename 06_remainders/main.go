package main

import "fmt"

func main() {
	x := 42 % 2
	fmt.Println(x)

	if x == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i, " = odd")
		} else {
			fmt.Println(i, " = even")

		}
	}
}
