package main

import "fmt"
import "math"

type Shaper interface {
	Area() float32
	Circumference() float32
}

type Rect struct {
	Width  float32
	Height float32
}

type Circle struct {
	Radius float32
}

func (r Rect) Area() float32 {
	return r.Width * r.Height
}

func (r Rect) Circumference() float32 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float32 {
	return math.Pi * 2 * c.Radius
}

func main() {
	r := Rect{10, 20}
	fmt.Printf("Rect w: %f, d: %f, Area: %f, Circumference: %f", r.Width, r.Height, r.Area(), r.Circumference())

	c := Circle{5}
	fmt.Printf("Circle r: %f, Area: %f, Circumference: %f", c.Radius, c.Area(), c.Circumference())

}
