package main

import "fmt"

func main() {
	visit([]int{2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})
}

func visit(numbers []int, callbacknij func(int)) {
	for _, n := range numbers {
		callbacknij(n)
	}
}
