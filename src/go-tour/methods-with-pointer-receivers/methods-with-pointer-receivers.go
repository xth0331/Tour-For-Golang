package main

import (
	"fmt"
	"math"
)

// Vertex is ..
type Vertex struct {
	X, Y float64
}

// Scale is ...
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs is ...
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v， Abs: %v\n", v, v.Abs())
}
