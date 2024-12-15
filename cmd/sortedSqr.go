package main

import (
	"math"
	"sort"
)

// TODO: Add TwoPointers
func sortedSquares(nums []int) []int {
	result := make([]int, 0)

	for _, v := range nums {
		result = append(result, int(math.Pow(float64(v), 2)))
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
