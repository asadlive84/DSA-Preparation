func minRemoveToMakeValid(s string) string {
	stack := []int{}
	index := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				continue
			} else if string(s[i]) == ")" && string(s[stack[len(stack)-1]]) == "(" {
				index[i] = true
				index[stack[len(stack)-1]] = true
				stack = stack[:len(stack)-1]
			}
		}
	}

	newStr := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == ")" {
			if index[i] {
				newStr += string(s[i])
			}
		} else {
			newStr += string(s[i])
		}
	}

	return newStr

}