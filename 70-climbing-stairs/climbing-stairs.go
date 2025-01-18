func helper(n int, memo map[int]int) int {
	if data, ok := memo[n]; ok {
		return data

	}
	if n < 0 {
		return 0
	}
	if n == 1 || n == 0 {
		return 1
	}

	memo[n] = helper(n-1, memo) + helper(n-2, memo)
	return memo[n]

}

func climbStairs(n int) int {
	memo := make(map[int]int)
	return helper(n, memo)
}