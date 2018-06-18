package main

import (
	"errors"
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
		return 0, errors.New("SQRT cannot be derived for negative numbers")
	}
	return 42, nil
}
