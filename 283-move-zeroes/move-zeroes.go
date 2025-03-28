func moveZeroes(nums []int) {
	left := 0
	right := 0

	for right < len(nums) {

		if nums[left] == 0 && nums[right] != 0 {
			nums[left] = nums[right]
			nums[right] = 0
		}

		if nums[left] != 0 {
			left++
		}

		right++
	}
}