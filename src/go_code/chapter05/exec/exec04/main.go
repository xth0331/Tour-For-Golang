package main

import (
	"fmt"
)

func main() {
	var age int
	var month int

	fmt.Println("请输入月份:")
	fmt.Scanln(&month)

	fmt.Println("年龄：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Println("旺季老人票：20")
		} else if age >= 18 {
			fmt.Println("旺季成人票：60")
		} else {
			fmt.Println("旺季儿童票：20")
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Println("淡季成人票：40")
		} else {
			fmt.Println("淡季半价票：20")
		}
	}
}
