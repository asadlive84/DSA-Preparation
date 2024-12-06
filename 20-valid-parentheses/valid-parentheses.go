func isValid(s string) bool {
	pMap := map[string]string{"}": "{", ")": "(", "]": "["}
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{" {
			stack = append(stack, string(s[i]))
		} else if string(s[i]) == ")" || string(s[i]) == "]" || string(s[i]) == "}" {
			if len(stack) == 0 {
				return false
			}
			if v, ok := pMap[string(s[i])]; ok {
				if stack[len(stack)-1] == v {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}

		}
	}

	return len(stack) == 0
}