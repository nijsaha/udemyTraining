package main

import "fmt"

func main() {
	age := 44
	fmt.Println("age = ", age)
	fmt.Println("&age = ", &age)

	changeme(&age)
	fmt.Println("new age = ", age)
	fmt.Println("new &age = ", &age)

	m := make([]string, 1, 25)
	fmt.Println(m)
	changeme1(m)
	fmt.Println(m)

}

//values gets passed.  For non reference types, address needs to be passed for changing original variable
func changeme(ageCopy *int) {
	*ageCopy = 20
	fmt.Println("ageCopy = ", ageCopy)
	fmt.Println("&ageCopy = ", &ageCopy)
	fmt.Println("*ageCopy = ", *ageCopy)
}

//values gets passed.  For reference types, address does not need to be passed.  As reference types are basically addresses anyways
//slices, maps and channels are reference types
func changeme1(z []string) {
	fmt.Println(z)
	z[0] = "Nij"
	fmt.Println(z)
}
