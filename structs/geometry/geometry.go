package geometry

import "math"

type Rectangle struct {
	Length int
	Width  int
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
