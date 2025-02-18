func helper(grid [][]int, r, c int, memo map[string]int) int {

	if data, ok := memo[fmt.Sprintf("%d-%d", r, c)]; ok {
		return data
	}

	if r == 0 && c == 0 {
		return grid[r][c]
	}
	if r < 0 || c < 0 {
		return math.MaxInt
	}

	memo[fmt.Sprintf("%d-%d", r, c)] = grid[r][c] + min(helper(grid, r, c-1, memo), helper(grid, r-1, c, memo))
	return memo[fmt.Sprintf("%d-%d", r, c)]
}

func minPathSum(grid [][]int) int {
	memo := make(map[string]int)
	r, c := len(grid)-1, len(grid[0])-1

	return helper(grid, r, c, memo)
}