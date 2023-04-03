package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {
	//结构体初始化
	person := Person{"zhagn", 18, "地址", 20}
	person = Person{name: "张三"}

	persons := []Person{
		{"zhagn", 18, "地址", 20},
		{"zhagn", 18, "地址", 17},
		{name: "张三"},
	}
	for k, v := range persons {
		fmt.Println(k, " - ", v.name)
	}

	for i := range persons {
		fmt.Println(persons[i])
	}

	//匿名结构体
	address := struct {
		province string
		city     string
		address  string
	}{
		"北京市",
		"昌平区",
		"xxx",
	}
	fmt.Println(address)

	fmt.Println(person)
}
