package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.side * s.side
}
func info(s shape) {
	fmt.Println("area = ", s.area())
}
func main() {
	c := circle{5}
	info(&c)
}
