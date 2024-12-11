func moveZeroes(nums []int) {

	for index, _ := range nums {
		for j := index + 1; j < len(nums); j++ {
			if nums[index] == 0 {
				nums[index], nums[j] = nums[j], nums[index]
			}
		}
	}

}