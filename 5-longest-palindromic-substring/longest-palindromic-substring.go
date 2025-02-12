func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	start, lastIndex := 0, 1
	table := make([][]int, n)

	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		table[i][i] = 1
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = 1
			start, lastIndex = i, i+2
		}

	}

	for length := 3; length <= n; length++ {

		for i := 0; i <= n-length; i++ {
			j := i + length - 1

			if s[i] == s[j] && table[i+1][j-1] == 1 {
				table[i][j] = 1
				start = i
				lastIndex = j + 1
			}
		}
	}

	//fmt.Println(table)
	//fmt.Println(start, lastIndex)

	return s[start:lastIndex]

}