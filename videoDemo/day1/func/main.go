package main

import "fmt"

type add_func func(int, int) int

func main() {
	//函数赋值
	fun := add
	sum := operator(fun, 100, 200)
	fmt.Println(sum)
	//改变基本类型的值
	a := 10
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}

func operator(ad add_func, a, b int) int {
	return ad(a, b)
}

func add(a, b int) int {
	return a + b
}

func modify(a *int) {
	*a = 100
}
