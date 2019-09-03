package main

import (
	"fmt"
)

// 输出一百以内的奇数，使用for+continue
func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
