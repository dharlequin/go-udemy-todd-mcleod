package main

import "fmt"

func main() {
	x := make([][]string, 2)

	x[0] = []string{"James", "Bond", "Shaken, not stirred"}
	x[1] = []string{"Miss", "Moneypenny", "Helloooooo, James."}

	for i, s := range x {
		fmt.Println(i, s)

		for _, v := range s {
			fmt.Printf("\t%v\n", v)
		}
	}
}
