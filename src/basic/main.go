package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	x, y := swap("jack", "rose")
	fmt.Printf("x: %s, y: %s \n", x, y)

	// 类型转换
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(u)
}

func swap(x, y string) (string, string) {
	return y, x
}
