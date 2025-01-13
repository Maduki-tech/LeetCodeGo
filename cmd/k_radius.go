package main

func getAverages(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, n)

	windowSize := 2*k + 1

	if windowSize > n {
		for i := range result {
			result[i] = -1
		}
		return result
	}

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	for i := 0; i < n; i++ {
		if i < k || i+k >= n {
			result[i] = -1
		} else {
			test1 := prefixSum[i+k+1]
			test2 := prefixSum[i-k]
			sum := test1 - test2
			result[i] = sum / windowSize
		}

	}

	return result
}
