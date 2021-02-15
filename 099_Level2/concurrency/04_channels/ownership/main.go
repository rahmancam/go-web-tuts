package main

import "fmt"

func main() {

	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 6; i++ {
				ch <- i
			}
		}()
		return ch
	}

	consumer := func(in <-chan int) {
		for v := range in {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Printf("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}
