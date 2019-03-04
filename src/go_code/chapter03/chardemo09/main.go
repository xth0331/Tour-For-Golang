package main

import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0' //字符0
	
	//直接输出byte值，直接输出ASCII码值。
	fmt.Println("c1 =",c1,"c2 =",c2)
	fmt.Printf("c1=%c c2=%c\n",c1,c2)   // 直接输出对应字符，格式化输出。

	// var c3 byte = '北' //overflow 溢出，范围问题。
	var c3 int = '北'
	fmt.Printf("c3 =%c c3对应码值=%d",c3,c3)
}