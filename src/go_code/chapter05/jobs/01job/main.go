package main

import (
	"fmt"
)

// 判断一个整数属于哪个范围  大于0， 小于0，等于0.

func main() {
	var num1 int64
	fmt.Println("请输入一个数：")
	fmt.Scanln(&num1)
	if num1 < 0 {
		fmt.Println("小于0")
	} else if num1 > 0 {
		fmt.Println("大于0")
	} else {
		fmt.Println("等于0")
	}
}
