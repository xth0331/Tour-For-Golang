package main

import (
	"fmt"
)

func main() {

	var score int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)

	// 多分支判断
	if score == 100 {
		fmt.Println("BWM")
	} else if score > 80 && score <= 99 {
		fmt.Println("iphone")
	} else if score > 60 && score <= 80 {
		fmt.Println("iPad")
	} else {
		fmt.Println("什么都么有")
	}

}
