package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}
type DoubleZero struct {
	Person
	LicenceToKill bool
}

func (p Person) greeting() {
	fmt.Println("Inner Person Greeting")
}

func (dz DoubleZero) greeting() {
	fmt.Println("Outer Double Zero Greeting")
}

func main() {
	p1 := Person{
		First: "Nij",
		Last:  "Saha",
		Age:   34,
	}
	p2 := DoubleZero{
		Person: Person{
			First: "Nayanika",
			Last:  "Bhaduri",
			Age:   33,
		},
		LicenceToKill: true,
	}

	p1.greeting()
	p2.greeting()
	p2.Person.greeting()
}
