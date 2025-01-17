func uniquePaths(row, col int) int {
	grid := gridConvert(row, col)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {

			curr := grid[row][col]

			if row+1 < len(grid) {
				grid[row+1][col] = grid[row+1][col] + curr
			}
			if col+1 < len(grid[0]) {
				grid[row][col+1] = grid[row][col+1] + curr
			}

		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func gridConvert(row, col int) [][]int {
	grid := make([][]int, row+1)
	for i:= range grid{
        grid[i]= make([]int, col+1)
    }
	grid[1][1] = 1
	return grid
}