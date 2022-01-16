package main

import "fmt"

func main() {
	x := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	x = append(x, 2048)
	fmt.Println(x)

	x = append(x, 4096, 4096*2, 4096*3)
	fmt.Println(x)

	y := []int{4096 * 4, 4096 * 5, 4096 * 6, 4096 * 7, 4096 * 8}
	x = append(x, y...)
	fmt.Println(x)
}
