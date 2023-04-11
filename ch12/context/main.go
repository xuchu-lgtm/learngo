package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func cpuIInfo(ctx context.Context) {
	defer wg.Done()
	fmt.Printf("tracid: %s\r\n", ctx.Value("key1"))

	for {
		select {
		case <-ctx.Done():
			fmt.Println("close")
			return
		default:
			fmt.Println("cpu info")
		}
	}
}

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	ctxValue := context.WithValue(ctx, "key1", "value1")
	go cpuIInfo(ctxValue)
	wg.Wait()
}
