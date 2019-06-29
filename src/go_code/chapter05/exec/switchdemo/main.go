package main

import (
	"fmt"
)

func main() {
	// 编写一个程序，该程序可以接受一个字符，比如:a,b,c,d,e,f,g
	// a表示星期一，b表示星期二... 根据用户输入显示对应数据。

	// 要求使用 switch 完成

	// 1.定义个变量接收字符
	// 2.使用switch判断
	var key byte
	fmt.Println("今天是周几，以a，b，c，d，e，f，g来表示周一到周日")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("输入有误")
	}
}
