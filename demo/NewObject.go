package main

import "C"
import (
	"fmt"
)

func main() {

	//a := Request{"1231","21312"}

	//创建一个空间为零的对象
	a := new(Request)
	fmt.Println(a)

	//创建长度为1的数组
	b := make([]int, 1)
	fmt.Println(b)
	//创建一个长度为零的数组
	c := new([]int)
	fmt.Println(c)

	//resp, err := http.Get("http://example.com/")
	//if err != nil {
	//
	//}
	//fmt.Println(resp)
}

type Request struct {
	Method string
	Params string
}
