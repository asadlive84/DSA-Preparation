func missingNumber(nums []int) int {
	totalNum := (len(nums) * (len(nums) + 1)) / 2

	totalSumInArr := 0
	for _, data := range nums {
		totalSumInArr += data
	}

	return totalNum - totalSumInArr
}