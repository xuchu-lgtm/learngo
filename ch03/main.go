package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	name := "go体系学习"
	bytes := []rune(name)
	fmt.Println(reflect.TypeOf(bytes))
	fmt.Println(len(bytes))

	for k, v := range bytes {
		fmt.Println(k, v)
	}

	username := "zhangsan"
	age := 18

	fmt.Printf("姓名：%s,年龄：%d \n", username, age)

	sprintf := fmt.Sprintf("姓名：%s,年龄：%d \n", username, age) //返回字符串结果
	fmt.Println(sprintf)

	//通过string的builder进行字符串拼接，高性能
	var builder = strings.Builder{}
	builder.WriteString(username)
	builder.WriteString(strconv.Itoa(age))

	result := builder.String()
	fmt.Println(result)
}
