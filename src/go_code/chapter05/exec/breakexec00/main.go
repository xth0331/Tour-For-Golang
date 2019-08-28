package main

import (
	"fmt"
)

func main() {
	// 100 以内数求和，求出当和第一次大于20的当前数。
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i // 求和
		if sum > 20 {
			fmt.Println("当前值为", i)
			break
		}
	}
}
