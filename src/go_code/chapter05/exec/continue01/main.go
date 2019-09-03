package main

import (
	"fmt"
)

// 从键盘读入不确定的数，判断正负，输入0结束退出。

var positiveCount int // 正数个数
var negativeCount int // 负数个数
var num int

func main() {
	for {
		fmt.Println("请输入:")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正数个数为%v\n", positiveCount)
	fmt.Printf("负数个数为%v\n", negativeCount)
}
