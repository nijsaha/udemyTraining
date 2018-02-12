package main

import (
	"fmt"
)

/*
Slices are reference types. referecing an underlying array.
3 main concepts to slices.  Referencing and Underlying array, lengths and Capacity
*/

func main() {
	//this is a slice.  Different from an array as number is not known.  Slice is basically a list
	//unlike like an array, slice lengths change at execution
	var a []int
	a = []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T - %v\n", a, a)
	fmt.Printf("%T - %v\n", b, b)

	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2]) //at index position
	//slicing a slice
	fmt.Println(mySlice[2:4]) //from index position 2 to 4 but not including 4
	//The above is also a index access.
	fmt.Println("Nirjhar Saha"[2])

	//slices have lengths invoked via len
	//slices have capacity invoked via cap
	//one can make a slice
	//make([]Type, len, capacity)
	/* - These two below have the same effect
	   make([]int, 50, 100)
	   new([100]int)[0:50]
	*/

	sl := make([]int, 0, 5) //if only one parameter is used, then both length and capacity becomes that number
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))

	for i := 0; i < 80; i++ {
		sl = append(sl, i)
		fmt.Println("Len: ", len(sl), " Cap: ", cap(sl), " Value: ", sl[i])
	}

	for i, values := range sl {
		fmt.Println(i, " - ", values)
	}

	greeting := []string{
		"Hello",
		"Bonjour",
		"Ohayo",
	}
	fmt.Println(greeting[:])   //all
	fmt.Println(greeting[1:])  //from position 1
	fmt.Println(greeting[:2])  //till before position 2 - note: 2 is excluded
	fmt.Println(greeting[1:2]) //from position 1 and till before position 2 - Note: 2 is excluded
	fmt.Println("Capacity of Greeting is ", cap(greeting))
	greeting = append(greeting, "Ola")

	for _, abc := range greeting {
		fmt.Println(abc)
	}
	fmt.Println("Capacity of Greeting is ", cap(greeting))

	dayofWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday"}
	greeting = append(greeting, dayofWeek...) //appending slices to another slice

	fmt.Println(greeting[:]) //all

	greeting = append(greeting[:2], greeting[3:]...) //deleting an element from a slice
	fmt.Println(greeting[:])                         //all

}
