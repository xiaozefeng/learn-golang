package main

import "fmt"

func main() {
	// 声明指针
	var p *int
	fmt.Println(p)

	// & 符号会生成一个指向其作用对象的指针。
	i := 42
	// 生成一个指针指向i
	p = &i

	// * 符号表示指针指向的底层的值
	// 通过指针 p 读取 i
	fmt.Println(*p)
	// 通过指针 p 设置 i
	*p = 21
	fmt.Println(i)
}
