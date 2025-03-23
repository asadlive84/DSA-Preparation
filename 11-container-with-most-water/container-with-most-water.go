func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	maxSum := 0
	for left < right {

		dist := right - left

		maxSum = max(min(height[left], height[right])*dist, maxSum)

		if height[left] <= height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		}

	}

	return maxSum
}