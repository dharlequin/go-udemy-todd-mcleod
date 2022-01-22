package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int32
	var iterations = 100
	var wg sync.WaitGroup

	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)

			fmt.Printf("\tcounter: %d\n", atomic.LoadInt32(&counter))

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
