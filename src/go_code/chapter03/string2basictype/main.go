package main

import (
	"fmt"
	"strconv"
)

// 演示golang中string转成基本数据类型

func main() {

	var str string = "true"
	var b bool
	// _ 忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "12345590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}
