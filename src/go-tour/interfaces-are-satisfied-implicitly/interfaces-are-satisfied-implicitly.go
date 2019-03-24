package main

import "fmt"

// I is ...
type I interface {
	M()
}

// T is ...
type T struct {
	S string
}

// M is ...此方法表示类型 T 实现了接口 I ，但我们无需显示声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
