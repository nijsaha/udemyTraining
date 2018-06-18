package main

import (
	"errors"
	"fmt"
	"log"
)

//ErrCustomErrorMsg gives a std custom error
var ErrCustomErrorMsg = errors.New("Sqrt cannot be derived of a negative number")

func main() {
	fmt.Printf("%T\n", ErrCustomErrorMsg)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(n int) (float64, error) {
	if n < 0 {
		return 0, ErrCustomErrorMsg
	}
	return 42, nil
}
