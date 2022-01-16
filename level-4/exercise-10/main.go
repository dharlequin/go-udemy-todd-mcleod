package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["q"] = []string{"Geek stuff", "Cooking", "Being precise"}

	delete(m, "no_dr")

	for key, val := range m {
		fmt.Println(key)
		for i, item := range val {
			fmt.Printf("\t%d\t%v\n", i, item)
		}
	}
}
