package main

import (
	"fmt"
)

// 声明一个函数

func test() bool {
	fmt.Println("test...")
	return true
}
// 短路与， && ：如果第一个条件为false，则第二个条件不会判断，最终结果为false。
// 短路或， || ： 如果第一个条件为true，则第二个条件不会判断，最终结果为true。
func main() {

	var i int = 10
//	if i < 9 && test() {
	if i > 9 || test() {
		fmt.Println("ok...")
	}
}
