package main

import "fmt"

func main() {
	x := [5]int{32, 64, 128, 256, 512}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
