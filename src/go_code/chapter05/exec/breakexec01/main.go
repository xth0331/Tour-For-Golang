package main

import (
	"fmt"
)

func main() {
	// 实现登录验证，有三次机会，如果用户名为张无忌，密码为888,提示登录成功，否则提示还有几次机会。
	var user string
	var password string
	for i := 2; i >= 0; i-- {
		fmt.Println("用户名：")
		fmt.Scanln(&user)
		fmt.Println("密码：")
		fmt.Scanln(&password)
		if user == "zwj" && password == "888" {
			fmt.Println("登录成功！")
			break
		}
		if i == 0 {
			fmt.Println("达到最大失败次数：请重新运行该程序！！！")
		} else {
			fmt.Printf("还有%d次机会。\n", i)
		}
	}
}
