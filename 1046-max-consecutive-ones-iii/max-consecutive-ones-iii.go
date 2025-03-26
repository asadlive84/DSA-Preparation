func longestOnes(nums []int, k int) int {

	left, right := 0, 0

	countZero := 0
	maxCount := 0

	for left <= right {

		if right >= len(nums) {
			break
		}

		if nums[right] == 0 {
			countZero++
		}

		for countZero > k {
			if nums[left] == 0 {
				countZero--
			}
			left++
		}

		count := right - left + 1

		if count > maxCount {
			maxCount = count
		}
		right++
	}

	return maxCount
}