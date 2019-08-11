package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 我们为了生成一个随机数，还需要个rand设置一个种子。
	// time.Now().Unix():返回一个从1970:01:01的0时0分0秒
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		// fmt.Println("n=", n)
		count++
		if n == 99 {
			break
		}
	}

	fmt.Println(count)
}
