package main

import (
	"fmt"
)

// 根据输入的月份和年份，求出该月天数。
// 如果闰年，2月是29 平年则是28
// 1、 3、 5、 7、 8、 10、 12 月是31
// 4 、6 、 9 、 11 月是30

func main() {
	var year int
	var month int
	fmt.Println("输入年份：")
	fmt.Scanln(&year)
	fmt.Println("输入月份：")
	fmt.Scanln(&month)
	// if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
	// 	switch month {
	// 	case 2:
	// 		fmt.Println("这个月有29天！")
	// 	default:
	// 		fmt.Println("请输入正确的年月日格式！")
	// 	}
	// } else {
	// 	switch month {
	// 	case 1, 3, 5, 7, 8, 10, 12:
	// 		fmt.Println("这个月有31天！")
	// 	case 4, 6, 9, 11:
	// 		fmt.Println("这个月有30天！")
	// 	case 2:
	// 		fmt.Println("这个月有28天！")
	// 	default:
	// 		fmt.Println("请输入正确的年月日格式！")
	// 	}

	// }

	switch month {
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			fmt.Println("这个月有29天！")
		} else {
			fmt.Println("这个月有28天！")
		}
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("这个月有31天！")
	case 4, 6, 9, 11:
		fmt.Println("这个月有30天！")
	default:
		fmt.Println("请输入正确的年月日格式！")
	}
}
