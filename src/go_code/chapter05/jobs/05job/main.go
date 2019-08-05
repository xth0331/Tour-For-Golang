package main

import (
	"fmt"
)

func main() {
	var user string
	var passwd string
	fmt.Println("用户名：")
	fmt.Scanln(&user)
	fmt.Println("密码")
	fmt.Scanln(&passwd)
	if user == "张三" && passwd == "1234" {
		fmt.Println("对了")
	} else {
		fmt.Println("错了")
	}
}
