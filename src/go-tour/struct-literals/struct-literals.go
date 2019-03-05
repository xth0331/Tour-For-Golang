package main

import "fmt"

type vertex struct {
	X, Y int
}

var (
	v1 = vertex{1, 2}  //创建一个 vertex类型的结构体
	v2 = vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = vertex{}      // X: 0 Y: 0
	p  = &vertex{1, 2} // 创建一个 *vertex类型的结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
