func largestAltitude(gain []int) int {

	prefix := make([]int, len(gain))

	prefix[0] = gain[0]

	max := prefix[0]

	for i := 1; i < len(gain); i++ {
		prefix[i] = prefix[i-1] + gain[i]
		if prefix[i] > max {
			max = prefix[i]
		}
	}
	
	if max < 0 {
		return 0
	}

	return max
}