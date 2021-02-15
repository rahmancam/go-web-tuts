package main

import (
	"fmt"
	"time"
)

func genMsg(out chan<- string) {
	defer close(out)
	for i := 0; i < 10; i++ {
		out <- fmt.Sprintf("Hello World %v", i)
		time.Sleep(time.Second * 1)
	}
}

func relayMsg(in <-chan string, out chan<- string) {
	defer close(out)
	for msg := range in {
		out <- msg
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	for msg := range ch2 {
		fmt.Println(msg)
	}
}
