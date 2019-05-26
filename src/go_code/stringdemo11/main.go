package main

import (
	"fmt"
)

// 演示golang中的string类型使用
func main() {
	// string的基本使用
	var address = "北京长城 110 hello world"
	fmt.Println(address)
	//``  原样输出
	var str2 = `
	package main

	import (
		"fmt"
	)

	// 演示golang中的string类型使用
	func main() {
		// string的基本使用
		var address  = "北京长城 110 hello world"
		fmt.Println(address)
	`
	fmt.Println(str2)
	// 字符串拼接方法
	var str = "hello " + "world"
	str += " haha!"
	fmt.Println(str)
	// 当一个拼接的操作很长， 可以分行写 "+" 保留在上一行
	var str4 = "hello " + "world" + "hello " +
		"world" + "hello " + "world"
	fmt.Println(str4)

	var a int          // 0
	var b float32      // 0
	var c float64      //0
	var isMarried bool // false
	var name string    // ""
	fmt.Printf("a=%d, b=%f, c=%f,isMarried=%v, name=%v", a, b, c, isMarried, name)
}
