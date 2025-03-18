func dailyTemperatures(temperatures []int) []int {
	record := make([]int, len(temperatures))
	stack := []int{}

	for i := range temperatures {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			record[prev] = i - prev

		}
		stack = append(stack, i)
	}
	return record
}