package main

import (
	"fmt"
)

//name in caps for exporting (i.e public)
type person struct {
	first string
	last  string
	age   int
}

//DoubleZero Name in caps for exported
type DoubleZero struct {
	person
	first         string
	licenceToKill bool
}

func main() {
	fmt.Println("Struct & embedded types")
	p1 := DoubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   30,
		},
		first:         "Double Zero Seven",
		licenceToKill: true,
	}
	p2 := DoubleZero{
		person: person{
			first: "Money",
			last:  "Peny",
			age:   20,
		},
		licenceToKill: false,
	}
	fmt.Println(p1.person, p1.first, p1.last, p1.licenceToKill, p1.person.first)
	fmt.Println(p1.person, p1.first, p1.last, p1.licenceToKill)
	fmt.Println(p2.person, p2.first, p2.last, p2.licenceToKill)

}
