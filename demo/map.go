package main

import "fmt"

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo     //变量声明
	personDB = make(map[string]PersonInfo) // make函数创建  myMap = make(map[string] PersonInfo, 100) 指定长度
	// delete(personDB, "12345")  元素删除
	// 往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	// 从这个map查找键为"12345"的信息
	c := personDB["12345"]
	fmt.Println(c.ID)
	// ok是一个返回的bool型，返回true表示找到了对应的数据
	person, ok := personDB["12345"]
	if ok { // 找到了
		fmt.Println(person)
		fmt.Println(a(ok))
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}

func a(x bool) int {
	if x { // 找到了
		return 1
	} else {
		return 0
	}
}
