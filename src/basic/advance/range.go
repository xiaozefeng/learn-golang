package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6}
	for _,v := range arr  {
		fmt.Println(v)
	}
}
