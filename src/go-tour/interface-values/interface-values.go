package main

import (
	"fmt"
	"math"
)

// I is ...
type I interface {
	M()
}

// T is ...
type T struct {
	S string
}

// M is ...
func (t *T) M() {
	fmt.Println(t.S)
}

// F is ...
type F float64

// M is ...
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
