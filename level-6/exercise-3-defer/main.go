package main

import "fmt"

func main() {
	defer foo()

	bar()
}

func foo() {
	fmt.Println("Function 1")
}

func bar() {
	fmt.Println("Function 2")
}
