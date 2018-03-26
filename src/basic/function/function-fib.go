package main

import "fmt"

func main() {
	// 使用闭包实现裴波那契函数
	f := fib()
	for i := 0; i < 10;i++{
		fmt.Println(f())
	}
}

func fib() func() int {
	count, a, b := 0, 1, 1
	return func() int {
		switch count {
		case 0, 1:
			count ++
		default:
			a, b = b, a+b
		}
		return b
	}
}
