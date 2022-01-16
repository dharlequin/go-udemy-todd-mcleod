package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "cosmic grey",
		},
		luxury: true,
	}

	fmt.Println(s.luxury)
	printCommonVals(s.vehicle)

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "crimson",
		},
		fourWheel: true,
	}

	fmt.Println(t.fourWheel)
	printCommonVals(t.vehicle)
}

func printCommonVals(v vehicle) {
	fmt.Printf("doors: %v, color: %v\n", v.doors, v.color)
}
