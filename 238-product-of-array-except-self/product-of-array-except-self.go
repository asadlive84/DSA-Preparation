func productExceptSelf(nums []int) []int {
	leftArr := make([]int, len(nums))
	leftArr[0] = 1
	first := 1
	for i := 1; i < len(nums); i++ {
		first = first * nums[i-1]
		leftArr[i] = first
	}

	rightArr := make([]int, len(nums))
	rightArr[len(nums)-1] = 1
	last := 1
	for i := len(nums) - 2; i >= 0; i-- {
		last = last * nums[i+1]
		rightArr[i] = last
	}

	for i, _ := range nums {
		nums[i] = leftArr[i] * rightArr[i]
	}
	return nums
}