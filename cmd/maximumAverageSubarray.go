package main

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	var left, right = 0, 0
	var sum float64
	var maxAvg float64 = math.Inf(-1)

	for right < len(nums) {
		var window = right - left + 1
		sum += float64(nums[right])

		if window == k {
			maxAvg = max(sum/float64(k), maxAvg)

			sum -= float64(nums[left])
			left++
		}
		right++
	}
	return maxAvg
}
