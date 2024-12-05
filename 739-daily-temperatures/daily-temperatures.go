func dailyTemperatures(temperatures []int) []int {

	stack := []int{}
	result := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		//fmt.Println("stack::", stack, " ", " i ", i, " ", temperatures[i], " result ", result)

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			stack = append(stack, i)
			result[i] = 0
			continue
		}

		if temperatures[i] == temperatures[stack[len(stack)-1]] {

			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
				result[i] = 0
				continue
			}
			result[i] = stack[len(stack)-1] - i
			stack = append(stack, i)
			continue
		}

		if temperatures[i] < temperatures[stack[len(stack)-1]] {
			result[i] = stack[len(stack)-1] - i
			stack = append(stack, i)
		}

	}

	return result
}