package main

import (
	"math"
	"fmt"
)

func main() {
	// 函数作为一个值
	fc := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(fc(3,4))
}
