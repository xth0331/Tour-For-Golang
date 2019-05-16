package main

import (
	"fmt"
)

func itob(i int) bool { return i != 0 }

func main() {
	ac := 1
	ad := 0
	var acc bool
	var ade bool
	acc = itob(ac)
	ade = itob(ad)
	fmt.Println(acc,ade)
}