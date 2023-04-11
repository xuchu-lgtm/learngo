package main

import (
	"fmt"
	"time"
)

func g1(ch chan struct{}) {
	time.Sleep(time.Second * 2)
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(time.Second * 3)
	ch <- struct{}{}
}

func main() {
	//select 类似于switch case 语句，但是select的功能和我们操作linux里面提供的io的select、poll、epoll
	//select 主要作用于多个channel

	//现在有个需求，我们现在有两个goroutine都在执行，但是我在主的goroutine

	g1Chan := make(chan struct{})
	g2Chan := make(chan struct{})
	go g1(g1Chan)
	go g2(g2Chan)

	// select 可以监控多个channel，任何一个channel返回都知道
	// 1. 某一个分支就绪了就执行该分支， 2、如果两个都就绪了，就是随机执行，目的是 为了防止饥饿
	time := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-g1Chan:
			fmt.Println("g1Chan")
		case <-g2Chan:
			fmt.Println("g2Chan")
		case <-time.C:
			fmt.Println("timeout")
			return
		}
	}
}
