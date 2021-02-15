package main

import "fmt"

func main() {
	demo3()
}

func demo3() {
	ch := make(chan int, 6) // buffered channels

	go func() {
		defer close(ch)

		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}

func demo2() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			ch <- i
		}
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

func demo1() {
	com := make(chan int)
	go func(a, b int) {
		defer close(com)
		c := a + b
		com <- c
	}(1, 2)
	value := <-com
	fmt.Printf("Computed value %v\n", value)
}
