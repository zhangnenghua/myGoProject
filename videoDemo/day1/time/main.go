package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now().UnixNano()
	//格式化时间 格式必须是2006/1/02 15:04这个时间
	now := time.Now()
	fmt.Println(now.Format("2006/1/02 15:04"))
	b := time.Now().UnixNano()
	fmt.Println((b - a) / 1000)
}
