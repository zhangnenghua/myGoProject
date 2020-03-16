package main

import "fmt"

func main() {
	i := 0
	switch i {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough //只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case;
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}
}
