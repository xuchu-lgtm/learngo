package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover if A", r)
		}
	}()

	var names map[string]string
	names["go"] = "go测试"

	return 0, errors.New("this is an error")
}

func main() {

	if a, err := A(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

}
