package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u *user) printUser() {
	fmt.Printf("Name: %v %v\n", u.First, u.Last)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Println("Frequent quotes:")

	for _, s := range u.Sayings {
		fmt.Printf("\t%v\n", s)
	}

	fmt.Println()
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	fmt.Println("SORTED BY AGE")
	sort.SliceStable(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	printUsers(users)

	fmt.Println("SORTED BY LAST")
	sort.SliceStable(users, func(i, j int) bool { return users[i].Last < users[j].Last })
	printUsers(users)

	fmt.Println("SORTED SAYINGS")
	for _, u := range users {
		sort.Strings(u.Sayings)
	}
	printUsers(users)
}

func printUsers(users []user) {
	for _, u := range users {
		u.printUser()
	}
}
