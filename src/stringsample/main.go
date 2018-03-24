package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World"
	// 是否包含Hello
	fmt.Println(strings.Contains(s, "Hello"))

	// 索引
	fmt.Println(strings.Index(s, "o"))

	ss := "1,2,3,4"
	// 切割
	splitWords := strings.Split(ss, ",")
	fmt.Println(splitWords)
	// 合并
	fmt.Println(strings.Join(splitWords, "-"))

	fmt.Println(strings.HasPrefix(ss, "1"))

	fmt.Println(strings.HasSuffix(ss, "4"))

}
