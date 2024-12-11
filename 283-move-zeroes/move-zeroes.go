func moveZeroes(nums []int) {

	i := 0
	for _, data := range nums {
		if data != 0 {
			nums[i] = data
			i++
		}

	}

	for i <= len(nums)-1 {
		nums[i] = 0
		i++
	}

}