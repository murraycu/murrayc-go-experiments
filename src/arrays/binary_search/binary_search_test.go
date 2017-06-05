package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var a = []int{0, 2, 4, 7, 9, 11, 15}
	if binary_search(a, 11) != 5 {
		t.Error("binary_search() returned unexpected result.")
	}
}
