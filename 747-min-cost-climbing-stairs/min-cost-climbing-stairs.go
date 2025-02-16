func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func helper(cost []int, n int, memo map[int]int) int {

	if data, ok := memo[n]; ok {
		return data
	}

	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		memo[n] = cost[n]
		return memo[n]
	}

	memo[n] = cost[n] + min(helper(cost, n-1, memo), helper(cost, n-2, memo))
	return memo[n]
}

func minCostClimbingStairs(cost []int) int {
	memo := make(map[int]int)
	return min(helper(cost, len(cost)-1, memo), helper(cost, len(cost)-2, memo))
}