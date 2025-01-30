func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(text1, text2 []rune, i, j int, memo [][]int) int {

	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if text1[i] == text2[j] {
		memo[i][j] = 1 + lcs(text1, text2, i+1, j+1, memo)
		//return memo[fmt.Sprintf("%d,%d", i, j)]
	} else {
		memo[i][j] = max(lcs(text1, text2, i+1, j, memo), lcs(text1, text2, i, j+1, memo))
	}
	return memo[i][j]

}

func longestCommonSubsequence(text1 string, text2 string) int {

	text1Rune := []rune(text1)
	text2Rune := []rune(text2)
	memo := make([][]int, len(text1Rune))

	for i, _ := range memo {
		memo[i] = make([]int, len(text2Rune))
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}

	return lcs(text1Rune, text2Rune, 0, 0, memo)
}