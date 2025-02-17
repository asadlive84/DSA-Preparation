func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(arr []int, n int, memo map[int]int) int {
	if data, ok := memo[n]; ok {
		return data
	}

	if n < 0 {
		return 0
	}
	memo[n] = max(arr[n]+helper(arr, n-2, memo), helper(arr, n-1, memo))
	return memo[n]
}

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1{
        return nums[0]

    }
	memo1 := make(map[int]int)
	memo2 := make(map[int]int)
	arr1 := nums[:len(nums)-1]
	arr2 := nums[1:]
	return max(helper(arr1, len(arr1)-1, memo1), helper(arr2, len(arr2)-1, memo2))

}