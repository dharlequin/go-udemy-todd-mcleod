package main

import "testing"

func TestSqrt(t *testing.T) {
	got, err := sqrt(10)
	if err != nil {
		t.Error("Unexpected error!")
	}

	if got != 42 {
		t.Errorf("sqrt(10) = %v; want 42", got)
	}
}
