package main

import (
	"fmt"
	"time"
)

//使用go 进行多线程进行处理程序

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		go Add(i, i)
	}
}
