func isSubsequence(s string, t string) bool {

	left := 0
	right := 0

	for left <= right {

		if left == len(s) {
			return true
		}

		if right >= len(t) {
			return false
		}

		if s[left] == t[right] {
			left++

		}
		right++

	}

	return false
}