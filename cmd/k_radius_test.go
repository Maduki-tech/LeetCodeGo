package main

import (
	"reflect"
	"testing"
)

func Test_getAverages(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want []int
	}{
		{"first", []int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3, []int{-1, -1, -1, 5, 4, 4, -1, -1 - 1}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAverages(tt.nums, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
