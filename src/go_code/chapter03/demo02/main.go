package main

import "fmt"

func main() {
	// golang的变量使用方式1
	// 第一种：指定变量类型，声明后若不赋值，使用默认值
	// int的默认值时0，其他类型默认值稍后
	var i int
	fmt.Println("i=", i)

	// 第二种：根据值自行判断变量类型（类型推倒）
	var num = 10.11
	fmt.Println("num=", num)
	// 第三种； 省略var，注意 :=左侧变量不应该是已经声明过的，否则会导致编译错误
	// 下面的方式等价于 var name srting name = "tom"
	// :不能省略，否则错误
	name := "tom"
	fmt.Println("name", name)
}
