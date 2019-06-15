package main

import (
	"fmt"
)
func main() {
	// 赋值运算符的使用演示
	// var i int
	// i = 10 //基本赋值

	// 有两个变量， a和b，要求将其进行交换，最终打印结果
	// a = 9 ， b = 2 ==> 2 =2 b=9
	a := 9
	b := 2
	fmt.Printf("交换前的情况是 a = %v, b = %v \n",a, b)
	// 定义一个临时变量
	t := a
	a = b
	b = t
	fmt.Printf("交换后的使用情况是a = %v, b = %v \n",a ,b)

	// 复合赋值操作
	a += 17  // a = a + 17
	fmt.Println("a=",a)
}