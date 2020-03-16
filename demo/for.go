package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(sum)
		sum += i
	}
	a := []int{1, 2, 3}
	for _, val := range a {
		fmt.Println(val)
	}
	for {
		sum++
		if sum == 100 {
			fmt.Println(sum)
			break
		}
	}

}
