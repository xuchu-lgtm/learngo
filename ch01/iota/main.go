package main

import "fmt"

func a() (int, bool) {
	return 0, false
}

var name = "test" //全局定义

func main() {

	//匿名变量，变量的作用域
	var _ int
	_, ok := a()
	if ok {
		//打印
	}

	var mainName = "main"
	fmt.Println(mainName)

	//iota，特殊常量，可以认为是一个可以被编译修改的常量
	const (
		ERR1 = iota + 1
		ERR2
		ERR25 = "ha" //iota内部仍然会增加计数器
		ERR3         //iota + 1
		ERR31        //iota + 1
		ERR32        //iota + 1
		ERR33        //iota + 1
		ERR34 = 100  //此处不会影响下面的数值，只对本身有作用
		ERR4  = iota
		ERR5
	)
	fmt.Println(ERR1, ERR2, ERR25, ERR3, ERR4, ERR5, ERR34)

	const (
		ERRNEW1 = iota //默认iota是 0 值
	)
	fmt.Println(ERRNEW1)

	{
		localName := "local"
		fmt.Println(localName)
	}

	/*
		如果中断了iota那么必须显式的恢复，后续会自动递增
		自增类型默认是int类型
		iota能简化const类型的定义
		每次出现const的时候，iota初始化为0
	*/
}
