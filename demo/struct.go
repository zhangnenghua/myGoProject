package main

import "fmt"

type rect struct {
	x, y          float64
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func main() {
	r := rect{1, 2, 3, 4}
	fmt.Println(r.area())
	rect1 := new(Rect)
	rect2 := &rect{}
	rect3 := &rect{0, 0, 100, 200}
	rect4 := &rect{width: 100, height: 200}
	fmt.Println(rect1, rect2, rect3, rect4)
}
