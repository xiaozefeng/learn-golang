package main

import "fmt"


type Animal struct {
	Name string
	Age int
}

var m map[string]Animal

var m1 = map[string]Animal{
	"rose": Animal{
		"rose", 21,
	},
	"lucy": Animal{
		"lucy", 20,
	},
}

var m2 = map[string]Animal{
	"rose": {"jack", 12},
	"lucy": {"lucy", 13},
}

func main() {
	m = make(map[string]Animal)
	m["jack"] = Animal{"jack", 18}
	fmt.Println(m)
	fmt.Println(m["jack"])

	fmt.Println("m1", m1)
	fmt.Println("m2", m2)

	// 检测某个值是否存在
	v, ok := m1["jack"]
	fmt.Println("The value:", v, "Present?", ok)

}
