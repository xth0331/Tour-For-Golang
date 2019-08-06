package main

import (
	"fmt"
)

// 判断考试成绩为什么等级，
// 60以下不及格
// 60-69 及格
// 70-79 良好
// 80-89 优良
// 90-100 优秀
func main() {
	var grade float64
	fmt.Println("输入分数：")
	fmt.Scanln(&grade)
	// switch {
	// case grade > 100:
	// 	fmt.Println("分数超过满分")
	// case grade >= 90 && grade <= 100:
	// 	fmt.Println("优秀")
	// case grade >= 80:
	// 	fmt.Println("优良")
	// case grade >= 70:
	// 	fmt.Println("良好")
	// case grade >= 60:
	// 	fmt.Println("及格")
	// case grade < 60:
	// 	fmt.Println("不及格")
	// }
	if grade > 100 {
		fmt.Println("分数超过满分")
	} else if grade >= 90 && grade <= 100 {
		fmt.Println("优秀")
	} else if grade >= 80 {
		fmt.Println("优良")
	} else if grade >= 70 {
		fmt.Println("良好")
	} else if grade >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
