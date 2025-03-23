func maxArea(height []int) int {

	left, right := 0, len(height)-1

	maxSum := 0

	for left < right {

		maxSum = max(min(height[right], height[left])*(right-left), maxSum)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}

	}

	return maxSum
}