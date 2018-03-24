package main

import (
	"encoding/xml"
	"fmt"
)

// 人物档案
type person struct {
	Name string `xml:"name,attr"`
	Age  int    `xml:"年龄,attr"`
}

func main() {
	// 转成成xml
	p := person{Name: "jack", Age: 18}

	var data []byte
	var err error

	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	p2 := new(person)
	if err := xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)
}
