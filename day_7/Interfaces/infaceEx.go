package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() float64 {
	return float64(r.Length) * float64(r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func init() {
	fmt.Println("Wellcome to the Interfaces")
}
