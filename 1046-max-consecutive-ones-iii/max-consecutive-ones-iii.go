func longestOnes(nums []int, k int) int {

	left := 0

	countZero := 0
	maxCount := 0

	for right:=0; right < len(nums); right++ {

		

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
		
	}

	return maxCount
}