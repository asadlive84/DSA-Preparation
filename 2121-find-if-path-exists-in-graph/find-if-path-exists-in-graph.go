func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	vertixs := convertEdge(n, edges)

	fmt.Println("vertixs: ", vertixs)

	visited := make(map[int]bool)

	stack := []int{source}
    
    visited[source] = true

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		

		stack = stack[:len(stack)-1]
    if current == destination {
			return true
		}
		

		for _, neighbour := range vertixs[current] {
			if !visited[neighbour] {
                visited[neighbour] = true
				stack = append(stack, neighbour)
			}
		}

	}

	return false
}

func convertEdge(n int, edges [][]int) map[int][]int {


	graph := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	return graph
}