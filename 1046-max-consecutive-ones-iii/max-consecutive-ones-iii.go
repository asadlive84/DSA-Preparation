func longestOnes(nums []int, k int) int {

	left, withZero, maxLength := 0, 0, 0

	for right := 0; right < len(nums); right++ {

		if nums[right] == 0 {
			withZero++
		}

		for withZero > k {
			if nums[left] == 0 {
				withZero--
			}
			left++
		}

		sum := right - left + 1
		if sum > maxLength {
			maxLength = sum
		}
	}

	return maxLength
}