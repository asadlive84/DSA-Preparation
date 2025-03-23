func isSubsequence(s string, t string) bool {
	left, right := 0, 0

	for right < len(t) {
		if left == len(s) {
			return true
		}
		
		if s[left] == t[right] {
			left++

		}
		right++
	}

	return left == len(s)
}