package main

import "fmt"

func main() {
	// new 返回的是一个指针  make 返回的是一个对象类型
	var p *[]int = new([]int)
	var v []int = make([]int, 10)
	fmt.Println(p)
	fmt.Println(v)
}
