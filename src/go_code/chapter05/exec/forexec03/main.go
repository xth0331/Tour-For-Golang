package main

import (
	"fmt"
)

// i 表示层数
// j 表示打印多少个星
// k 表示*前面的空格数

func main() {
	var laynum int
	fmt.Println("输入层数：")
	fmt.Scanln(&laynum)
	for i := 1; i <= laynum; i++ {
		for k := 1; k <= laynum-i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	for i := laynum-1; i >= 1; i-- {
		for k := 1; k <= laynum-i; k++ {
			fmt.Printf(" ")
		}
		for j := 2*i - 1; j >= 1; j-- {
			if j == 1 || j == 2*i-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()

	}
}
