package main
import "fmt"


// 演示golang中整数类型使用
func main() {

	var i int = 1
	fmt.Println("i =",i)

	//测试一下int8的范围 -128 -127
	var j int8 = -128
	fmt.Println("j =",j)

	//测试一下 uint8的范围
	var k uint16 = 255
	fmt.Println("k =",k)

	//int ,uint(), rune(), byte(0-255)
	var a = -1
	fmt.Println("a = ",a)
	var b uint = 255
	fmt.Println("b =",b)
	var c rune = 2147483647
	fmt.Println("c =",c)
	var d byte = 255
	fmt.Println("d =",d)


}