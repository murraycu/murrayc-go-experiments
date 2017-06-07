package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var a = []int{0, 2, 4, 7, 9, 11, 15}
	pos, found := binary_search(a, 11)

	if !found {
		t.Error("binary_search() didn't find the value.")
	}

	if pos != 5 {
		t.Error("binary_search() returned unexpected position.")
	}
}
