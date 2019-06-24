package main

import (
	"fmt"
)

func main() {
	var grade float64
	var gender int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&grade)

	if grade <= 8.0 {
		fmt.Println("请输入性别(1代表男，2代表女):")
		fmt.Scanln(&gender)
		if gender == 1 {
			fmt.Println("进入男子组决赛")
		} else {
			fmt.Println("女子组决赛")
		}
	} else {
		fmt.Println("淘汰")
	}
}
