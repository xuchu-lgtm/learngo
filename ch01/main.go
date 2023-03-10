package main

import "fmt"

// 全局变量和局部变量
var name = "test"
var age = 10
var ok bool

var (
	name1 = "test"
	ages1 = 10
)

func main() {
	//go是静态语言，静态语言和动态语言相比变量差异很大
	/*
		1.变量必须先定于后使用
		2.变量必须有类型
		3.类型定下来后不能改变
	*/
	//定义变量的方式
	//var name int

	age := 1
	//go 语言中变量变量定义不使用是不行的
	fmt.Println(age)

	//2. 多变量定义
	var user1, user2, user3 = "test", "test2", "test3"
	fmt.Println(user1, user2, user3)

	/*
		注意：
			变量必须先定义再使用
			go语言是静态语言，要求变量的类型和赋值类型一致
			变量名不能冲突（局部变量优先级比较高）
			简洁变量定义不能用于全局变量
			变量是有零值的
			定义了变量一定要使用，否则会报错
	*/
}
