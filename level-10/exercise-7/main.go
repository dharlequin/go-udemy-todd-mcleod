package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {

		go func() {
			addNumber(c)
		}()
	}

	for v := 0; v < 100; v++ {
		fmt.Println(<-c)
	}
}

func addNumber(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
