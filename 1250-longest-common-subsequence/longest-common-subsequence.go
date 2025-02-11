func longestCommonSubsequence(fStr, lStr string) int {

	totalRow := len(lStr)
	totalCol := len(fStr)

	table := make([][]int, totalRow+1)

	for index, _ := range table {
		table[index] = make([]int, totalCol+1)
	}

	for r := 1; r <= totalRow; r++ {
		for c := 1; c <= totalCol; c++ {
			if lStr[r-1] == fStr[c-1] {
				table[r][c] = 1 + table[r-1][c-1]
			} else {
				max := table[r-1][c]
				if table[r][c-1] > max {
					max = table[r][c-1]
				}
				table[r][c] = max
			}
		}
	}

	//fmt.Println(table)

	return table[totalRow][totalCol]
}