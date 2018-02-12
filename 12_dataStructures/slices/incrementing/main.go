package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	mySlice[0] = 7
	fmt.Println(mySlice[:])
	fmt.Println(mySlice[0])
	mySlice[0]++ //value 8 gets stored in mySlice[0] similar to mySlice[0]=mySlice[0]+1
	//mySlice[0]=mySlice[0]+1
	//mySlice[0] += 1
	fmt.Println(mySlice[:])
	fmt.Println(mySlice[0])
	mySlice[0] = 9
	fmt.Println(mySlice[:])
	fmt.Println(mySlice[0])

}
