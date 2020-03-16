package main

import "fmt"

func main() {
	defer fmt.Println("defer")

	//多参数传递函数
	sum := sum(10, 10)
	fmt.Println(sum)

	result := concat("hello ", "world ", "znh")
	fmt.Println(result)
}

func sum(a int, arg ...int) int {
	sum := a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func concat(a string, arg ...string) (result string) {
	result = a
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return
}
