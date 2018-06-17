package main

import (
	"fmt"
	"strconv"
)

func main() {

	//conversion - casting on the left side with the type
	var x = "12"
	//y := 9
	z, _ := strconv.Atoi(x)
	fmt.Println(z)

	//assertions for interfaces - casting on the right side with a .(type)
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)

	} else {
		fmt.Printf("Value is not a string\n")
	}

}
