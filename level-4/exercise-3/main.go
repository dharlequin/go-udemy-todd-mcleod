package main

import "fmt"

func main() {
	x := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	a := x[:5]
	b := x[5:]
	c := x[2:7]
	d := x[1:6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
