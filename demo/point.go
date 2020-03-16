package main

import "fmt"

//&x指针代表指向变量的内存地址   *x代表指针指向地址的值

func PointTransmit() {
	var temp = 1
	fmt.Println("point transmit before:", temp)
	method2(&temp)
	fmt.Println("point transmit after", temp)
}
func method2(temp *int) {
	fmt.Println("method2 before:", *temp)
	*temp = 2
	fmt.Println("method2 after:", *temp)
}

//结果
//point transmit before: 1
//method2 before: 1
//method2 after 2
//point transmit after: 2

//从上面的代码中可以看出指针传递后， 对指针所指向地址的值修改后，对其它方法是可见的
