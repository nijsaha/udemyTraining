package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side

}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())

}
func main() {
	s := square{10}
	c := circle{10}
	info(s)
	info(c)
}
