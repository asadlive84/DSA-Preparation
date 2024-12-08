func strToInt(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return i
}

func evalRPN(tokens []string) int {
	stack := []string{}

	for i := 0; i < len(tokens); i++ {

		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			newSum := 0
			if tokens[i] == "+" {
				newSum = strToInt(n2) + strToInt(n1)
			}
			if tokens[i] == "-" {
				newSum = strToInt(n2) - strToInt(n1)
			}
			if tokens[i] == "*" {
				newSum = strToInt(n2) * strToInt(n1)
			}
			if tokens[i] == "/" {
				newSum = strToInt(n2) / strToInt(n1)

			}
			stack = append(stack, strconv.Itoa(newSum))

		} else {
			stack = append(stack, tokens[i])

		}
	}
	if len(stack) == 0 {
		return 0
	}

	return strToInt(stack[0])

}