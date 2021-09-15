package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	// defer close(c)
	go send(c)
	go receive(c)

	for {
	}
}

func send(c chan int) {
	defer func() {
		close(c)
		fmt.Println("send end")
	}()

	i := 0
	for {
		fmt.Println("ready send")
		select {
		case c <- i:
			i++
			if i > 20 {
				return
			}
			time.Sleep(time.Duration(200) * time.Millisecond)
		case <-time.After(1 * time.Second):
			fmt.Println("send 1s pass")
		}
	}
}

func receive(c chan int) {
	defer func() {
		fmt.Println("receive end")
	}()

	for {
		select {
		case v, ok := <-c:
			if !ok {
				fmt.Println("channel is closed")
				return
			}
			fmt.Println(v)
		case <-time.After(1 * time.Second):
			fmt.Println("receive 1s pass")
		}
	}
}
