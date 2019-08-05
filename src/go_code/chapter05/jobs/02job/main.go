package main

import (
	"fmt"
)

// 判断一个年份是否为闰年

func main() {
	var year int
	fmt.Println("输入年:")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("是闰年！")
	} else {
		fmt.Println("不是闰年！")
	}
}
