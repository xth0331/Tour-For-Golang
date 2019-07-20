package main

import (
	"fmt"
)

func main() {
	var laynum int
	fmt.Println("输入层数：")
	fmt.Scanln(&laynum)
	for i := 1; i <= laynum; i++ {
		for k := 1; k <= laynum-i; k++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
