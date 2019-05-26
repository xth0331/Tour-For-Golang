package main

import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0' //字符0

	//直接输出byte值，直接输出ASCII码值。
	fmt.Println("c1 =", c1, "c2 =", c2)
	fmt.Printf("c1=%c c2=%c\n", c1, c2) // 直接输出对应字符，格式化输出。

	// var c3 byte = '北' //overflow 溢出，范围问题。
	var c3 int = '北'
	fmt.Printf("c3 =%c c3对应码值=%d\n", c3, c3)

	var c4 = 22269 // 22269 -> '国'
	fmt.Printf("c4=%c\n", c4)
	// 字符类型是可以进行运算的， 相当于一个整数，运算时是按照码值运算。
	var n1 = 10 + 'a' // 10 + 97 = 107
	fmt.Println("n1=", n1)
}
