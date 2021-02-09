package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("# of Go routines running :", runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancelFn := context.WithCancel(ctx)

	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println("Running # : ", n)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("Stil working :", n)
					time.Sleep(100 * time.Millisecond)
				}
				fmt.Println("Go routines running : ", runtime.NumGoroutine())
			}
		}(i)
	}

	time.Sleep(5 * time.Second)
	cancelFn()
	fmt.Println(time.Second)
	fmt.Println("# of Go routines running :", runtime.NumGoroutine())

}
