package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
	circumferece() float32
}
type square struct {
	length float32
}
type circle struct {
	radius float32
}

func (s square) area() float32 {
	return s.length * s.length
}
func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (s square) circumferece() float32 {
	return s.length * 4
}
func (c circle) circumferece() float32 {
	return 2 * math.Pi * c.radius
}
func printArea(l shape) {
	fmt.Printf("area of %T is %v \n", l, l.area())
	fmt.Printf("area of %T is %v \n", l, l.circumferece())
}
func main() {
	s1 := square{20}
	c1 := circle{20.0}
	printArea(s1)
	printArea(c1)
}
