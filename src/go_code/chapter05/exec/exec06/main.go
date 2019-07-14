package main

import (
	"fmt"
)

func main() {
	var b int = 0
	var c = a + b
	for a:=0; a<=6; {
		if a+b == 6 {
			fmt.Println(a, "+", b, "=", c)
		}
		a++
	}
}
