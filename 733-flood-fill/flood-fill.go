func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	orginalColor := image[sr][sc]
	if orginalColor == color {
		return image
	}

	fill := func(image [][]int, row, col, color, orginalColor int) {
		stack := [][]int{{row, col}}

		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			r, c := curr[0], curr[1]

			if r < 0 || c < 0 || r >= len(image) || c >= len(image[0]) || image[r][c] == color {
				continue
			}
			if image[r][c] != orginalColor {
				continue

			}
			image[r][c] = color

			stack = append(stack, []int{r + 1, c})
			stack = append(stack, []int{r - 1, c})
			stack = append(stack, []int{r, c + 1})
			stack = append(stack, []int{r, c - 1})

		}

	}

	fill(image, sr, sc, color, orginalColor)

	return image
}