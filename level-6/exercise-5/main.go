package main

import (
	"fmt"
	"math"
)

type square struct {
	length float32
}

func (s square) area() float32 {
	return s.length * s.length
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float32
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	square := square{
		length: 2,
	}

	info(square)

	circle := circle{
		radius: 3,
	}

	info(circle)
}
