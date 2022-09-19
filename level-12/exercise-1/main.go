package main

import (
	"fmt"
	dog "go/learning-with-todd/level-12/exercise-1/dog"
)

func main() {
	myDogYears := dog.Years(12)
	fmt.Printf("My Dog is %v human years old", myDogYears)
}
