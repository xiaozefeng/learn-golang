package main

import "fmt"

func main() {
	p := Person{"jack", 18}
	fmt.Println(p)
	p.Name = "rose"
	p.Age = 19

	fmt.Println(p)
	// 通过间接指针来修改
	c := &p
	c.Name = "lucy"
	c.Age = 20
	fmt.Println(*c)
}

type Person struct {
	Name string
	Age int
}
