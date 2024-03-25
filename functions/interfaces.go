package functions

import (
	"fmt"
	"math"
)

func Interfaces() {
	c := circle{
		radius: 1.45,
	}

	s := square{
		length: 46,
		width:  89,
	}

	fmt.Println("this is the circle area", c.area())
	fmt.Println("this is the square area", s.area())
	fmt.Println(info(c))
	fmt.Println(info(s))
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	a := s.length * s.width
	return a
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) string {
	return fmt.Sprintln("this is the figure area", s.area())
}
