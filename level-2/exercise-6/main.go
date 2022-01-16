package main

import "fmt"

const (
	x = 2017
)

const (
	a = x + iota
	b = x + iota
	c = x + iota
	d = x + iota
)

func main() {
	fmt.Printf("%d\t%d\t%d\t%d", a, b, c, d)
}
