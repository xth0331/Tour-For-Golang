package main

import (
	"fmt"
)

func main() {
	// 演示 & 和 *
	a := 100
	fmt.Println("a的地址 =", &a)

	var ptr *int = &a
	fmt.Println("ptr 指向的值是 =", *ptr)
}
