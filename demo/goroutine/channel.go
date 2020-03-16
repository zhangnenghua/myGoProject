package main

import (
	"fmt"
)

func count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(chs[i])
	}
	for _, ch := range chs {
		//更优雅的加锁方式 只有等待channel有值才能正常返回
		fmt.Println(<-ch)
	}
}
