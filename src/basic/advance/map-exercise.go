package main

import (
	"strings"
	"fmt"
)

/*
 统计字符串中字符出现的次数
 */
func main() {
	arr := strings.Split("asfhakshgkqwnflkqhl", "")
	wordCount := make(map[string]int)
	fmt.Println(arr)
	for _, v := range arr {
		_, ok := wordCount[v]
		if ok {
			wordCount[v] += 1
		} else {
			wordCount[v] = 1
		}
	}
	fmt.Println(wordCount)
}
