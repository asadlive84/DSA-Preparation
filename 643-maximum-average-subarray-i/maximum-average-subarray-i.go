func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for i := 0; i < k; i++ {
		sum = sum + nums[i]
	}
	//left, right := 0, k-1
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum

		}
	}

	return float64(maxSum) / float64(k)
}