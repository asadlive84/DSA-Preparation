func gridTraveller(row, col int, memo map[string]int) int {
	if data, ok := memo[fmt.Sprintf("%d, %d", row, col)]; ok {
		return data
	}

	if row == 1 && col == 1 {
		return 1
	}
	if row == 0 || col == 0 {
		return 0
	}

	memo[fmt.Sprintf("%d, %d", row, col)] = gridTraveller(row-1, col, memo) + gridTraveller(row, col-1, memo)
	return memo[fmt.Sprintf("%d, %d", row, col)]
}

func uniquePaths(m int, n int) int {
    memo := make(map[string]int)
    return gridTraveller(m,n, memo)
}