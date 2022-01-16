package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	alias     string
}

func main() {
	p := person{
		firstName: "Peter",
		lastName:  "Quill",
		alias:     "Star-Lord",
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)
}

func changeMe(p *person) {
	p.alias = "Scut-Lord"
}
