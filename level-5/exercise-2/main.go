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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, val := range m {
		printValues(val)
	}
}

func printValues(p person) {
	fmt.Printf("%v %v\n", p.first, p.last)

	for _, flavor := range p.icFlavors {
		fmt.Printf("\t%v\n", flavor)
	}

	fmt.Println()
}
