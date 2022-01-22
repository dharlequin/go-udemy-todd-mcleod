package main

import "fmt"

func main() {
	x := "Scope of main"

	{
		x := "Enclosed scope"
		fmt.Println(x)
	}

	fmt.Println(x)
}
