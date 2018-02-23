package main

import (
	"fmt"
)

//name in caps for exporting (i.e public)
type Person struct {
	first string
	last  string
	age   int
}

//Name in caps for exported
type DoubleZero struct {
	Person
	licenceToKill bool
}

func main() {
	fmt.Println("Struct & embedded types")
	p1 := DoubleZero{
		Person: Person{
			first: "James",
			last:  "Bond",
			age:   30,
		},
		licenceToKill: true,
	}
	p2 := DoubleZero{
		Person: Person{
			first: "Money",
			last:  "Peny",
			age:   20,
		},
		licenceToKill: false,
	}
	fmt.Println(p1.Person, p1.first, p1.last, p1.licenceToKill)
	fmt.Println(p2.Person, p2.first, p2.last, p2.licenceToKill)

}
