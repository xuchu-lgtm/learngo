package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	p1 := Person{"zha", 19}
	p2 := Person{"zha", 18}

	//在golang中结构体是可以判断是否相等的
	if p1 == p2 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
