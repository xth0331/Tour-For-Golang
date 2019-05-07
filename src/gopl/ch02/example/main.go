package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	fmt.Println(f, err)
}
