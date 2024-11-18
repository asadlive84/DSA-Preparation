func removeElement(nums []int, val int) int {
    numsLen := len(nums)
	newIndex := 0
	for i := 0; i < numsLen; i++ {

		if nums[i] != val {
			nums[newIndex] = nums[i]
			newIndex++
		}
	}
	nums = nums[:newIndex]
	fmt.Printf("nums %+v \n", nums)
	return len(nums)

}