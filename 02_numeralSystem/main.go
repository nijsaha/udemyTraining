package main

import "fmt"

func main() {
	fmt.Println(42)
	fmt.Printf("%d - %b \n", 42, 42)   //binary
	fmt.Printf("%d - %o \n", 42, 42)   //octa
	fmt.Printf("%d - %#o \n", 42, 42)  //symbolic octa
	fmt.Printf("%d - %x \n", 42, 42)   //hex
	fmt.Printf("%d - %#x \n", 42, 42)  //symbolic hex
	fmt.Printf("%d - %X\n", 42, 42)    //HEX
	fmt.Printf("%d - %#X\n", 42, 42)   //symbolic HEX
	fmt.Printf("%d - %U\n", 42, 42)    //Unicode
	fmt.Printf("%d - \t%#U\n", 42, 42) //symbolic Unicode

	for i := 60; i < 90; i++ {
		fmt.Printf("%d \t %b \t %o \t %x \t %q \n", i, i, i, i, i) //symbolic Unicode
	}
}
