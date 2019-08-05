package main

import (
	"fmt"
)

func main() {
	var num2 int
	var b int
	var s int
	var g int
	fmt.Println("输入一个数，判断是否为水仙数：")
	b = num2 / 100
	s = num2 % 100 / 10
	g = num2 % 100 % 10
	fmt.Println(b, s, g)
	if b*b*b+s*s*s+g*g*g == num2 {
		fmt.Println("是")
	} else {
		fmt.Println("不是")
	}
}
