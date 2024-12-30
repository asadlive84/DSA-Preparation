func findDuplicate(nums []int) int {
	newNumsMaps := make(map[int]bool)

	for _, data := range nums {
		if _, ok := newNumsMaps[data]; ok {
			return data
		}
		newNumsMaps[data] = true
	}
	return 0
}