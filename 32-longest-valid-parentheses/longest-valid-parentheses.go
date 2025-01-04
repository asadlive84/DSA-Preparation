type Paren struct {
	Data  rune
	Index int
}

func longestValidParentheses(s string) int {

	dataRune := []rune(s)
	stack := []Paren{}
	n := make([]int, len(s), len(s))
	fmt.Println(":n:", n, len(dataRune))
	for i, data := range dataRune {
		if data == ')' {
			if len(stack) > 0 && stack[len(stack)-1].Data == '(' {
				n[i] = 1
				n[stack[len(stack)-1].Index] = 1
				stack = stack[:len(stack)-1]
			} else {
				continue
			}
		} else if data == '(' {
			stack = append(stack, Paren{Data: data, Index: i})

		}
	}

	count := 0
	max := 0
	for _, d := range n {
		if d == 1 {
			count++
		} else {
			count = 0
		}

		if count > max {
			max = count
		}
	}

	return max
}