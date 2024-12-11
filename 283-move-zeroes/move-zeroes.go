func moveZeroes(nums []int) {

	newNums := make([]int, len(nums), len(nums))
	i := 0
	for _, data := range nums {

		if data != 0 {
			newNums[i] = data
			i++
		}

	}

	for index, data := range newNums {
		nums[index] = data
	}

}