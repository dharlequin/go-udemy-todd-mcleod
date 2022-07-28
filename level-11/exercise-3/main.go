package main

import "fmt"

type customErr struct {
	err string
}

func (ce customErr) Error() string {
	return ce.err
}

func main() {
	ce := customErr{
		err: "custom error",
	}

	foo(ce)
}

func foo(e error) {
	fmt.Println(e.Error())
}
