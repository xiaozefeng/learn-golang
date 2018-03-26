package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
/**
	没有类，也可以把方法加到结构体上
 */
func main() {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
