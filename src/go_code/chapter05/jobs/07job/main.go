package main

import (
	"fmt"
	"math"
)

// 根据公式（身高-108） *2 = 体重，可以有十斤左右的浮动。来观察测试者体重是否合适。

func main() {
	var height float64
	var weight float64
	fmt.Println("输入身高：")
	fmt.Scanln(&height)
	fmt.Println("输入体重：")
	fmt.Scanln(&weight)
	if math.Abs((height-108)*2-weight) <= 10 {
		fmt.Println("标准")
	} else {
		fmt.Println("不标准")
	}
}
