package main

import (
	"fmt"
)

// 演示golang中的指针类型

func main() {

	// 基本数据类型在内存的布局
	var i int = 10
	// i 的地址是什么， &i
	fmt.Println("i的地址=", &i)
	// 1. ptr 是一个指针变量
	// 2. ptr 的类型是 *int
	// 3. ptr 本身的值是&i
	var ptr *int = &i  // *int 声明指针类型变量
	fmt.Printf("ptr=%v\n", ptr)   
	fmt.Printf("ptr 的地址=%v\n", &ptr)
	fmt.Printf("ptr 指向的值=%v\n", *ptr)
}
