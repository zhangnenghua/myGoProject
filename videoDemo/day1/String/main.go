package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串替换
	str := "hello world"
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println(result) //hello you

	//字符串计数
	count := strings.Count(str, "l")
	fmt.Println(count) //3

	//以空格分割成一个slice
	str2 := "aa bb cc dd"
	result2 := strings.Fields(str2)
	fmt.Println(result2) // [aa bb cc ...]

	//自定义分割符
	result3 := strings.Split(str2, ",")
	fmt.Println(result3)

	//strings.Itoa(i int)：把一个整数i转成字符串
	//strings.Atoi(str string)(int, error)：把一个字符串转成整数
}

func main1() {
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
