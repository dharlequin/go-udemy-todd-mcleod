package main

import "fmt"

func main() {
	x := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
