package main

import "fmt"

func main() {
	f := foo()

	f()
}

func foo() func() {
	return func() {
		fmt.Println("Func in a func")
	}
}
