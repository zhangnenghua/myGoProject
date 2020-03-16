package main

import "fmt"

//匿名函数
func main() {
	fmt.Println(test(100, 200))
	multiplication()
}

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}(a, b)
	return result
}

func multiplication() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
