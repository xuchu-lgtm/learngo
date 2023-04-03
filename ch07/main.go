package main

import (
	"fmt"
	"strconv"
)

type MyInt int

func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))
}

func main() {

	type MyInt = int

	var i MyInt = 1
	fmt.Println(i)

	var a interface{} = "text"

	switch a.(type) {
	case string:
		fmt.Println("this is string")
	}

}
