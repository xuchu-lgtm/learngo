package main

import (
	"fmt"
	"time"
)

func main() {

	var msg chan string

	msg = make(chan string, 0)

	/*	msg <- "hello"

		data := <-msg

		fmt.Println(data)*/

	go func(msg chan string) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all done")
	}(msg)

	msg <- "hello"
	msg <- "hello2"

	close(msg)

	time.Sleep(time.Second * 10)

	var c1 chan string
	var c2 chan<- string // 只能放
	var c3 <-chan string // 只能取

	c1 <- "test"
	c2 <- "23e"
	data := <-c3
	fmt.Println(data)
}
