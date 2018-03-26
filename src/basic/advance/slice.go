package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("p==", arr)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("index: %d, value:%d \n", i, arr[i])
	}

	// 切片
	fmt.Println("arr[1:4]", arr[1:4])
	fmt.Println("arr[0:3]", arr[:3])
	fmt.Println("arr[3:]", arr[3:])

}
