func removeDuplicates(nums []int) int {
	numberMap := make(map[int]bool)
	writeIndex := 0

	for _, data := range nums {
		if !numberMap[data] {
			numberMap[data] = true
			nums[writeIndex] = data
			writeIndex++
		}
	}

	nums = nums[:writeIndex]

	return len(nums)
}