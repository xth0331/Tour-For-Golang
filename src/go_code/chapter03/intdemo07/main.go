package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//整型的使用细节
	var n1 = 100 //? n1是什么类型
	// 这里介绍如何查看某个变量的数据类型
	// fmt.printf() 可以用作格式化输出。
	fmt.Printf("n1的数据类型是:%T \n", n1)

	// 如何在程序查看某个变量的占用字节大小和数据类型
	var n2 int64 = 10000000000000
	// unsafe.Sizeof(n1) 是unsafe包的函数，可以返回变量的占用的字节数。
	fmt.Printf("n2的数据类型是: %T,占用的字节数是: %d\n", n2, unsafe.Sizeof(n2))
}
