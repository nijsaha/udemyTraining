package main

import (
	"fmt"
	"log"
)

func main() {
	x, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)

	} else {
		fmt.Println(x)
	}

}
func sqrt(n int) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("SQRT cannot be derived for negative numbers %v", n) //little bit more context
	}
	return 42, nil
}
