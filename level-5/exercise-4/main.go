package main

import "fmt"

func main() {
	person := struct {
		relatives      map[string]string
		visitedPlanets []string
	}{
		relatives: map[string]string{
			"mother": "Maureen Robinson",
			"father": "John Robinson",
		},
		visitedPlanets: []string{"Unknown Planet 1", "Unknonw Planet 2", "Alpha Centauri"},
	}

	fmt.Println(person.relatives)
	fmt.Println(person.visitedPlanets)
}
