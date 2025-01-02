func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image // No need to fill if the target color is the same as the original.
	}

	rows, cols := len(image), len(image[0])
	stack := [][]int{{sr, sc}} // Start with the initial position.

	for len(stack) > 0 {
		// Pop the last element from the stack.
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		r, c := curr[0], curr[1]

		// Ensure the current cell is within bounds and matches the original color.
		if r >= 0 && r < rows && c >= 0 && c < cols && image[r][c] == originalColor {
			// Update the cell with the new color.
			image[r][c] = color

			// Add neighboring cells to the stack.
			stack = append(stack, []int{r + 1, c}, []int{r - 1, c}, []int{r, c + 1}, []int{r, c - 1})
		}
	}

	return image
}