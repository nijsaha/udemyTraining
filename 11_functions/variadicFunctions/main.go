package main

import "fmt"

func main() {
	fmt.Println(average(23, 54, 67, 89, 100))
	data := []float64{2, 4}
	fmt.Println(average(data...)) //using variadic arguments
	//for arguments, dots are in the back

	greatest := max(2, 6, 9, 1, 34, 89, 23)
	fmt.Println(greatest)
}

//using variadic parameters
//for parameters, dots are in the front
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total = total + v

	}
	return total / float64(len(sf))
}

func max(numbers ...int) int {
	fmt.Printf("%T\n", numbers)
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}
