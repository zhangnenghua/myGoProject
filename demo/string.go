package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "baidu.com"

	fmt.Println(str)

	str1 := strings.HasPrefix(str, "http://")
	if !str1 {
		str = fmt.Sprintf("http://%s", str)
	}
	fmt.Println(str)

	str2 := strings.HasSuffix(str, "/")
	if !str2 {
		str = fmt.Sprintf("%s/", str)
	}
	fmt.Println(str)
}
