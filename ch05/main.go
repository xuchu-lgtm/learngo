package main

import "fmt"

func main() {
	/*	var courses1 [3]string
		var courses2 [4]string

		courses1[0] = "go"
		courses1[1] = "grpc"
		courses1[2] = "gin"

		fmt.Printf("%T \r\n", courses1)
		fmt.Printf("%T \n", courses2)
	*/

	/*courses1 := [3]string{"go", "grpc", "gin"}
	for _, value := range courses1 {
		fmt.Println(value)
	}*/

	courses2 := [3]string{2: "gin"}
	for _, value := range courses2 {
		fmt.Println(value)
	}

	courses3 := [...]string{"a", "b"}
	fmt.Printf("%T \r\n", courses3)

	for _, row := range courses3 {
		fmt.Println(row)
	}

}
