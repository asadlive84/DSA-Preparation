func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeAlternately(word1 string, word2 string) string {

	maxLen := max(len(word1), len(word2))
	newStr := ""
	for i := 0; i < maxLen; i++ {

		if i < len(word1) {
			newStr += string(word1[i])
		}

		if i < len(word2) {
			newStr += string(word2[i])
		}

	}

	return newStr
}