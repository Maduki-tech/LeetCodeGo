package main

import (
	"reflect"
	"testing"
)

func TestSortedSqr(t *testing.T) {

	input := []int{-4, -1, 0, 3, 10}
	expected := []int{0, 1, 9, 16, 100}
	result := sortedSquares(input)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v got %v", expected, result)
	}

}
