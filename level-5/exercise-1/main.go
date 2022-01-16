package main

import "fmt"

type person struct {
	first     string
	last      string
	icFlavors []string
}

func main() {
	p1 := person{
		first:     "Don",
		last:      "West",
		icFlavors: []string{"chocolate", "hazelnut"},
	}

	p2 := person{
		first:     "Penny",
		last:      "Robinson",
		icFlavors: []string{"pistachio", "strawberry"},
	}

	printFlavors(p1)
	printFlavors(p2)
}

func printFlavors(person person) {
	for _, flavor := range person.icFlavors {
		fmt.Println(flavor)
	}
}
