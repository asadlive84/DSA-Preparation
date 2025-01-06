func lengthOfLongestSubstring(s string) int {
	strRune := []rune(s)
	wordMap := make(map[string]int)

	left := 0
	right := 0

	maxCount := 0
	for left <= right {

		if right >= len(strRune) {
			break
		}

		if _, ok := wordMap[string(strRune[right])]; !ok {
			wordMap[string(strRune[right])] = right
		} else {
			if wordMap[string(strRune[right])] >= left {
				left = wordMap[string(strRune[right])] + 1
			}

			wordMap[string(strRune[right])] = right
		}

		count := right - left + 1
		if count > maxCount {
			maxCount = count
		}

		right++
	}

	//fmt.Println(wordMap)
	return maxCount

}