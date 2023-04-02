package main

import "fmt"

func main() {
	//var c byte //主要适用于存放字符，存放字符串
	var c2 rune //也是字符
	c2 = 'a'
	fmt.Printf("c=%c", c2)

	var name string
	name = "i am zhang"
	fmt.Println(name)
}
