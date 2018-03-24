package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整形转换成字符串
	fmt.Println(strconv.Itoa(10))

	// 字符串转换成整形
	fmt.Println(strconv.Atoi("123"))

	// 转换boolean
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("F"))
	fmt.Println(strconv.ParseBool("T"))

	// 转换float
	fmt.Println(strconv.ParseFloat("3.14", 32))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	// format
	fmt.Println(strconv.FormatBool(true))

	// 进制转换
	// 2进制
	fmt.Println("二进制", strconv.FormatInt(123, 2))
	fmt.Println("四进制", strconv.FormatInt(123, 4))
	fmt.Println("8进制", strconv.FormatInt(123, 8))
	fmt.Println("10进制", strconv.FormatInt(123, 10))
	fmt.Println("16进制", strconv.FormatInt(123, 16))

}
