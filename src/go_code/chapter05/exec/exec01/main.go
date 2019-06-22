package main

import (
	"fmt"
)

func main() {
	//编写程序，声明2个int32型变量并赋值。判断两个数之和，如果大于等于50，打印"hello world！"

	var n1 int32 = 9
	var n2 int32 = 40
	if n1+n2 >= 50 {
		fmt.Println("hello world!")
	} else {
		fmt.Println("xiaoyu50")
	}

	// 声明2个float64变量并赋值，判断第一个大于10.0
	// 且第2个数小于20.0. 打印两个数之和
	var n3 float64 = 7.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20 {
		fmt.Println(n3 + n4)
	}
}
