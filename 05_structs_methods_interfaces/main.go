package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	base   float64
	height float64
}

type Circle struct {
	radius float64
}

func (triangle Triangle) Perimeter() float64 {
	return 2 * (triangle.base + triangle.height)
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.base * triangle.height
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}
