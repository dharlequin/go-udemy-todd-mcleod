package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			addNumber(c)
		}()
	}

	wg.Wait()

	for v := range c {
		fmt.Println(v)
	}
}

func addNumber(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
