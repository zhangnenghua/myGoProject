package main

import "fmt"

type BASE struct {
	Name string
}

func (base *BASE) Foo() {
	fmt.Println(1)
}
func (base *BASE) Bar() {
	fmt.Println(base.Name)
}

//在“派生类”Foo没有改写“基类”Base的成员方法时，相应的方法就被“继承”，例如在 例子中，调用foo.Foo()和调用foo.Base.Foo()效果一致。
type Foo struct {
	base BASE
}

func (foo *Foo) Bar() {
	foo.base.Bar()
	foo.Bar()
}
