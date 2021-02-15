package main

import (
	"fmt"
	"sync"
)

func main() {
	demo2()
}

func demo2() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x)
		}(i)
	}

	wg.Wait()
}

func demo1() {
	var wg sync.WaitGroup
	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i : %v\n", i)
		}()
		fmt.Println("return from function")
		return
	}
	incr(&wg)
	wg.Wait()
	fmt.Println("Done!")

}
