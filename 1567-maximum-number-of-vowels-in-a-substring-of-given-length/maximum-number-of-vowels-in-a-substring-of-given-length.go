func maxVowels(s string, k int) int {

	//byteStr := make([]byte, 0, k)

	vowelsMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	count := 0

	for i := 0; i < k; i++ {
		if vowelsMap[s[i]] {
			count++
		}
	}

	maxCount := count

	for i := k; i < len(s); i++ {
		if vowelsMap[s[i-k]] {
			count--
		}

		if vowelsMap[s[i]] {
			count++
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount

}