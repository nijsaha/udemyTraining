package main

import (
	"fmt"
)

//multi dimensional slices
func main() {
	fmt.Println("multi dimensional slices")
	records := make([][]string, 0) //[] holds the slice of []string.
	//student1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	records = append(records, student1)
	fmt.Println(student1[:])
	fmt.Println(records[:][:])

	student2 := make([]string, 4)
	student2[0] = "Nij"
	student2[1] = "Saha"
	student2[2] = "80.00"
	student2[3] = "75.00"

	records = append(records, student2)
	fmt.Println(student2[:])
	fmt.Println(records)

	//3 ways to create a slices - Shorthand, var and make. Make i sthe best way to do slices
	shorthandSlice := []int{}
	var varSlice []int
	//varSlice[0] = 1
	makeSlice := make([]int, 3)
	fmt.Println(shorthandSlice == nil)
	fmt.Println(varSlice == nil)
	fmt.Println(varSlice)

	fmt.Println(makeSlice == nil)

}
