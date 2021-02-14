package main

import (
	"fmt"
	"time"
)

func hello(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	// Direct call
	hello("Direct call")

	go hello("go-routine 1")

	go func() {
		hello("Anonymous")
	}()

	test := hello
	go test("Function Value call")

	// wait for goroutines to end
	time.Sleep(time.Second * 1)

	fmt.Println("Done!")
}
