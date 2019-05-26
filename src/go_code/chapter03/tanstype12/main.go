package main

import (
	"fmt"
)

// 演示golang中基本数据类型的转换
func main() {
	var i int32 = 100
	// 希望将 i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i) // 低精度 -> 高精度

	fmt.Printf("i=%v n1=%v n2=%v n3=%v", i, n1, n2, n3)
}
