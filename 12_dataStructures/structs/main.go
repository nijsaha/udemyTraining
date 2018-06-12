package main

import "fmt"

type myAge int //can create your own type
type person struct {
	first string
	last  string
	age   myAge
}

//func (receiver) functionName(function arguments) returnType
func (p person) fullname() string {
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Howard", "Lenoard", 30}
	fmt.Println(p1, " - ", p1.first, ", ", p1.last, ", ", p1.age)
	fmt.Println(p2, " - ", p2.first, ", ", p2.last, ", ", p2.age)
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())

}
