type Value struct {
	Data  string
	Index int
}

func minRemoveToMakeValid(data string) string {
	stack := []Value{}
	runeData := []rune(data)

	for index, d := range runeData {
		if d == ')' {
			if len(stack) > 0 && stack[len(stack)-1].Data == "(" {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, Value{Data: string(d), Index: index})
			}

		} else if d == '(' {
			stack = append(stack, Value{Data: string(d), Index: index})
		}
	}

	removeIndices := make(map[int]bool)

	for _, data := range stack {
		removeIndices[data.Index] = true
	}
	info := []rune{}
	for index, _ := range runeData {
		if !removeIndices[index] {
			info = append(info, runeData[index])
		}
	}

	return string(info)
}