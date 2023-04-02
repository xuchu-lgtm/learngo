package main

import (
	"fmt"
)

func add2(a, b int) int {
	return a + b
}

func add(items ...int) (sum int, err error) {
	for _, item := range items {
		sum += item
	}
	return sum, nil
}

func cal(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return nil
	}
}

func add3(a, b int) {
	fmt.Printf("sum is: %d\r\n", a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func auoIncrement() func() int {
	local := 0 //一个函数中访问另一个函数的局部变量 不行的 闭包，函数可以随时归零
	return func() int {
		local += 1
		return local
	}
}

func main() {
	//go函数支持普通函数、匿名函数、闭包
	/*
		go中函数是“一等公民”
		1. 函数本身可以当做变量
		2. 匿名函数 闭包
		3. 函数可以满足接口
	*/

	sum, ok := add(1, 2, 4, 6)
	if ok != nil {
		fmt.Println("error")
	}
	fmt.Println(sum)

	funcVal := add

	sum, ok = funcVal(1, 2, 3)
	fmt.Println(sum)

	cal("+")()

	addFunc := func(a, b int) {
		fmt.Printf("sum is: %d \r\n", a+b)
	}

	callback(1, addFunc)

	//闭包
	nextFunc := auoIncrement()
	for i := 0; i < 10; i++ {
		fmt.Println(nextFunc())
	}

}
