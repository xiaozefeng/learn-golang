package main

import "fmt"

func main() {
	// ch <- v    // 将 v 送入 channel ch。
	// v := <-ch  // 从 ch 接收，并且赋值给 v。
	// （“箭头”就是数据流的方向。）
	// 和 map 与 slice 一样，channel 使用前必须创建：
	a := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)
	fmt.Println(ch)
	x, y := <-ch, <-ch // 从通道中获取
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println(x + y)

}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将sum 送人c
}
