package main

import (
	"fmt"
)

func main() {
	var laynum int 1
	for h := 0 ; h <= laynum; h++{
		for i:=1 ; i<= h ;i++ {
			fmt.Println("*")
		}
	}
}