package main

import "fmt"

func main() {
	// map 是一个Key(索引)和value(值)的无序集合，主要是查询方便 o(1)
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}

	fmt.Println(courseMap["go"])

	courseMap["mysql"] = "mysql入门"

	fmt.Println(courseMap)

	for k, v := range courseMap {
		fmt.Println(k, v)
	}

	if value, ok := courseMap["net"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println(value)
	}

	delete(courseMap, "go")

	//map不是线程安全的
	fmt.Println(courseMap)

}
