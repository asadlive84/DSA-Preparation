func dfs(grid [][]byte, row, col int) {
	stack := [][]int{{row, col}}

	for len(stack) > 0 {
		//fmt.Println("stack: ", stack)
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		r, c := curr[0], curr[1]

		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '0' {
			continue
		}

		grid[r][c] = '0'

		stack = append(stack, []int{r + 1, c})
		stack = append(stack, []int{r - 1, c})
		stack = append(stack, []int{r, c + 1})
		stack = append(stack, []int{r, c - 1})
	}
}

func numIslands(grid [][]byte) int {
	islandCount := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				islandCount++
				dfs(grid, row, col)
			}
		}
	}

	return islandCount

}