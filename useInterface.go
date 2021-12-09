package main

import (
	"fmt"
	"math"
	"github.com/fuadsuleyman/interface/shape"
)

type square struct {
	X float64
}
type circle struct {
	R float64
}

// implement the shape interface for the square type
func (s square) Area() float64 {
	return s.X * s.X
}
func (s square) Perimeter() float64 {
	return 4 * s.X
}

// implement the shape interface for the circle type
func (s circle) Area() float64 {
	return s.R * s.R * math.Pi
}
func (s circle) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

func Calculate(x shape.Shape) {
	val, ok := x.(circle)
	if ok {
		fmt.Printf("Is a circle and radius is %v\n", val.R)
	}
	v, ok := x.(square)
	if ok {
		fmt.Printf("Is a square and side is %v\n", v.X)
	}
	fmt.Println("Area:", x.Area())
	fmt.Println("Perimeter:", x.Perimeter())
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := circle{R: 5}
	Calculate(y)
}
