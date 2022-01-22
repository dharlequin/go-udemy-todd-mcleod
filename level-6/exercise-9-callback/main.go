package main

import "fmt"

func main() {
	bar(func() {
		fmt.Println("print stuff")
	})
}

func bar(f func()) {
	f()
}
