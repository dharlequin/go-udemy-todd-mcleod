package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Routines", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("I'm a Rat!")

		wg.Done()
	}()

	go func() {
		fmt.Println("I'm a Mermaid!")

		wg.Done()
	}()

	fmt.Println("Routines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Routines", runtime.NumGoroutine())
}
