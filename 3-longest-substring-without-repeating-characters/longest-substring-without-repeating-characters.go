func lengthOfLongestSubstring(str string) int {

	stringMap := make(map[string]int)

	left := 0
	right := 0
	maxlength := 0
	for left <= right {
		if right >= len(str) {
			break
		}

		if position, ok := stringMap[string(str[right])]; !ok {
			stringMap[string(str[right])] = right

		} else {
			if stringMap[string(str[right])] >= left {
				left = position + 1
			}

			stringMap[string(str[right])] = right
		}

		currentLength := right - left + 1
		if currentLength > maxlength {
			maxlength = currentLength
		}
		right++

	}

	return maxlength

}