package main

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		// TODO: Add test cases.
		{"1", "Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
