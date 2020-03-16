package main

import "fmt"

type admin struct {
	name string
	age  int
}

func main() {
	e, f := returns(1, 2)
	fmt.Println(e)
	if f != (admin{}) {
		fmt.Println(f)
	}
	base := new(Base)
	person(base)

}

func returns(a, b int) (c int, ad admin) {
	ad.age = 1
	ad.name = "213"
	return a + b, ad
}

func person(base *Base) {
	fmt.Println(base)
}

type Base struct {
	name string
}
