package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

//DoubleZero :caps means exported
type DoubleZero struct {
	person
	LicenceToKill bool
}

func (p person) greeting() {
	fmt.Println("Inner Person Greeting")
}

func (dz DoubleZero) greeting() {
	fmt.Println("Outer Double Zero Greeting")
}

func main() {
	p1 := person{
		First: "Nij",
		Last:  "Saha",
		Age:   34,
	}
	p2 := DoubleZero{
		person: person{
			First: "Nayanika",
			Last:  "Bhaduri",
			Age:   33,
		},
		LicenceToKill: true,
	}

	p1.greeting()
	p2.greeting()
	p2.person.greeting()
}
