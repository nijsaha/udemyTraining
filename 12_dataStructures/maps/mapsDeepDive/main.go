package main

import (
	"fmt"
)

//maps within a map - multi dimensional maps
func main() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	fmt.Printf("%T\n", elements)
	fmt.Println(elements)
	fmt.Println("Length: ", len(elements))

	//range loop for Maps
	for key, val := range elements {
		fmt.Println(key, " - ", val)
	}

}
