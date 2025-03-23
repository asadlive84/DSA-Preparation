func maxOperations(nums []int, k int) int {
    slices.Sort(nums)
	left := 0
	right := len(nums) - 1

	count := 0

	for left < right {
		sum := (nums[left] + nums[right])
		if sum == k {

			count++
			left++
			right--
			continue
		}

		if sum > k {
			right--

		} else {
			left++
		}
	}

	return count

}