package main

import (
	"fmt"
)
// golang小数使用
func main() {
	
	var price float32 = 1.11
	fmt.Println("price =",price)
	var num1 float32 = -0.00089
	var num2  = -78009656.09
	fmt.Println("num1 =",num1 , "num2 =",num2)

	// 尾数部分可能丢失，造成精度损失，-123.0000901
	var num3 float32 = -123.0000901
	var num4 = -123.0000901
	fmt.Println("num3 =",num3,"num4 =",num4)
	fmt.Printf("num4的数据类型是%T\n",num4)
	
	// 默认声明 float64
	var num5 = 1.1
	fmt.Printf("num5的数据类型是 %T\n",num5)
	//十进制数形式：如5.12 .512 (小数点不能省略)
	num6 := 5.12
	num7 := .123 // 0.123
	fmt.Println("num6 =",num6, "num7 =",num7)

	//科学计数法形式
	num8 := 5.1234e2  //等于5.1234 * 10的2次方
	num9 := 5.1234E2  //跟上面相等
	fmt.Println("num8 =",num8,"num9 =",num9)
}
