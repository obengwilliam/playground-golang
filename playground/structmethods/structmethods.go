package structmethods

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Calulate the perimeter of a rectangle
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Calulate the perimeter of a rectangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}
