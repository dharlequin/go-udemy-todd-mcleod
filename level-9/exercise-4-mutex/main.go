package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter = 0
	var iterations = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v

			fmt.Printf("\tcounter: %d\n", counter)
			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
