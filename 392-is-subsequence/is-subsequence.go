func isSubsequence(s string, t string) bool {
	left, right := 0, 0

	for right <= len(t) {
		if left == len(s) {
			return true
		}
		
		if right <len(t) && s[left] == t[right] {
			left++

		}
		right++
	}

	return false
}