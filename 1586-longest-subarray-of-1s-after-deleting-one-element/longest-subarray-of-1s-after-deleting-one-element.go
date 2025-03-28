func longestSubarray(nums []int) int {

	left, maxLength, zero := 0, 0, 0

	for right := 0; right < len(nums); right++ {

		if nums[right] == 0 {
			zero++
		}

		for zero >= 2 {
			if nums[left] == 0 {
				zero--
			}

			left++
		}

		currLength := right - left + 1 - zero

		if currLength > maxLength {
			maxLength = currLength
		}
	}

	if zero == 0 {
		return maxLength - 1
	}

	return maxLength

}