package main

import (
	"fmt"
)

// 这里演示一下指定标签的形式来使用 break
func main() {
	lable1:
	for i := 0; i < 4; i++ {
		// lable: 设置一个标签
		for j := 0; j <= 10; j++ {
			if j == 2 {
				// break 默认会跳出最近的for循环
				break lable1
				// break
			}
			fmt.Println("j=", j)
		}
	}
}
