package main

import (
	"sort"
	"fmt"
)

func main() {
	arr := []int{3, 6, 8, 9, 10, 0, 1, 2}
	sort.Ints(arr)

	for index, element := range arr {
		fmt.Printf("index= %d, element=%d \n", index, element)
	}
}
