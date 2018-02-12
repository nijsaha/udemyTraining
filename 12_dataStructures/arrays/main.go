package main

import (
	"fmt"
)

func main() {
	//different data structues in go
	/*
	   arrays -numbers in the brackets [2]
	   slices - no numbers in the brackets []
	   maps
	   struct
	*/

	var a []string  //slice
	var b [5]string //array
	fmt.Println("data structures")
	fmt.Println(a)
	fmt.Println(b)

	//ascii
	var x [58]string
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)

	}
	fmt.Println(len(x))
	fmt.Println(x)

	//UTF-8
	var y [256]int
	for j := 0; j < 256; j++ {
		y[j] = j

	}
	for _, v := range y {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}

}
