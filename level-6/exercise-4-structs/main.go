package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hey there! My name is %v %v. I'm %d.", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Thomas",
		last:  "Anderson",
		age:   33,
	}

	p.speak()
}
