package main

import "fmt"

type Consumer struct {
	ch chan string
}

type Broker struct {
	consumer []*Consumer
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumer = append(b.consumer, c)
}

func (b *Broker) Produce(msg string) {
	for _, c := range b.consumer {
		c.ch <- msg
	}
}

func main() {
	b := &Broker{consumer: make([]*Consumer, 0, 10)}
	c1 := &Consumer{ch: make(chan string, 1)}
	c2 := &Consumer{ch: make(chan string, 1)}

	b.Subscribe(c1)
	b.Subscribe(c2)

	b.Produce("hello")

	fmt.Println(<-c1.ch)
	fmt.Println(<-c2.ch)
}
