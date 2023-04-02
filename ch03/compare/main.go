package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "hello"
	b := "bello"
	fmt.Println(a != b)
	fmt.Println(a > b)

	name := "这 this is - 工程师go,go"
	fmt.Println(strings.Contains(name, "go"))

	fmt.Println(strings.Count(name, "go"))

	fmt.Println(strings.Split(name, "-"))

	fmt.Println(strings.HasPrefix(name, "这"))
	fmt.Println(strings.HasSuffix(name, "工程师"))

	fmt.Println(strings.Index(name, "go"))

	fmt.Println(strings.Replace(name, "go", "net", 2))

	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("java"))

	fmt.Println(strings.Trim("#$hello #go#", "#$"))
}
