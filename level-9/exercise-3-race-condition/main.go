package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var counter = 0
	var iterations = 100
	var wg sync.WaitGroup

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v

			fmt.Printf("\tcounter: %d\n", counter)

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
