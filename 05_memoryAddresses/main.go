package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a's memory address is ", &a)
	fmt.Printf("%d", &a)

	//pointers
	var b = &a //referencing the memory address
	fmt.Println("Pointers")
	fmt.Println(b)
	fmt.Println(*b) //derefencing - give the value

	*b = 42
	fmt.Println(a)

	var meters float64
	const metersToYards float64 = 1.09361

	fmt.Print("Enter meters swam")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println("meters is", meters)
	fmt.Println("yards is", yards)

	x := 5
	fmt.Printf("%p\n", &x)
	zero(&x)

	fmt.Println(x)
}

func zero(x *int) {
	fmt.Printf("%p\n", x)
	*x = 0
}
