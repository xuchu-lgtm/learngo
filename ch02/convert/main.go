package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 12

	/*a1 := 12
	var a3 int
	a3 = 13*/

	//var b = uint8(a)
	//var c = int32(a3)
	var f64 = float64(a)
	fmt.Println(f64)

	//类型要求很严格
	type IT int
	var c IT = IT(a)
	fmt.Println(c)

	//字符串转数字
	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	var myi = 32
	fmt.Println(strconv.Itoa(myi))

	// 字符串转换为float32,转换为bool
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	parseInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	// 0 = false ， 1 = true 字符串转bool
	parseBool, err := strconv.ParseBool("0")
	if err != nil {
		fmt.Println("parseBool error")
	}
	fmt.Println(parseBool)

	//bool类型转字符串
	formatBool := strconv.FormatBool(false)
	fmt.Println(formatBool)

	formatFloat := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Println(formatFloat)

	formatInt := strconv.FormatInt(42, 16)
	fmt.Println(formatInt)
}
