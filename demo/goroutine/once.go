package main

import (
	"fmt"
	"sync"
)

var a string
var once sync.Once

func setup() {
	a += "hello, world"
}
func doprint() {
	//全局唯一性操作
	once.Do(setup)
	//setup()
	println(a)
}
func main() {
	request := &Request{}
	request.Method = "123"
	fmt.Println(request.Method)
	go doprint()
	go doprint()
	go doprint()
	go doprint()
	go doprint()
	go doprint()
	go doprint()
}

type Request struct {
	Method string "method"
	Params string "params"
}
