package main

import (
	"fmt"
	"time"
)

func main() {
	demo3()
}

func demo1() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println(m1)
		case m2 := <-ch2:
			fmt.Println(m2)
		}
	}

}

func demo2() {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "one"
	}()

	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}

func demo3() {

	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("message %v", i)
		}
	}()

	// for i := 0; i < 2; i++ {
	// 	m := <-ch
	// 	fmt.Println(m)

	// 	fmt.Println("Processing..")
	// 	time.Sleep(1500 * time.Millisecond)
	// }

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no msg received")
		}

		fmt.Println("Processing..")
		time.Sleep(1500 * time.Millisecond)
	}
}
