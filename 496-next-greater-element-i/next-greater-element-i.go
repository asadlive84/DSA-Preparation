func nextGreaterElement(nums1 []int, nums2 []int) []int {
	resMap := make(map[int]int)
	res := make([]int, len(nums2))
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			resMap[nums2[i]] = -1
			res[i] = -1
			stack = append(stack, nums2[i])
		}

		if nums2[i] < stack[len(stack)-1] {
			res[i] = stack[len(stack)-1]
			resMap[nums2[i]] = stack[len(stack)-1]
			stack = append(stack, nums2[i])
			continue
		}

		for len(stack) != 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			if nums2[i] < stack[len(stack)-1] {
				res[i] = stack[len(stack)-1]
				resMap[nums2[i]] = stack[len(stack)-1]
				stack = append(stack, nums2[i])
				continue
			}
		}

		if len(stack) == 0 {
			res[i] = -1
			resMap[nums2[i]] = -1
			stack = append(stack, nums2[i])
		}
	}
	result := make([]int, len(nums1))
	for index, data := range nums1 {
		if _, ok := resMap[data]; ok {
			result[index] = resMap[data]
		}
	}

	return result
}