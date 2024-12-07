func longestValidParentheses(s string) int {
	//count := 0
	stack := []int{}
	idex := make([]int, len(s))
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				continue
			} else if s[i] == ')' && s[stack[len(stack)-1]] == '(' {
				idex[i] = 1
				idex[stack[len(stack)-1]] = 1
				stack = stack[:len(stack)-1]
			}
		}
	}
	//fmt.Println("idex ::", idex)
	max := 0
	for i := 0; i < len(s); i++ {
		currentVal := 0
		for j := i; j < len(s); j++ {
			if idex[j] == 0 {
				break
			}
			currentVal += idex[j]

		}

		if currentVal > max {
			max = currentVal
		}
	}

	return max
}