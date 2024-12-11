func majorityElement(nums []int) int {
	newMap := make(map[int]int)

	for _, key := range nums {
		if data, ok := newMap[key]; ok {

			newMap[key] = data + 1

		} else {
			newMap[key] = 1
		}

	}
	max := 0
	key := 0

	for k, v := range newMap {

		if v > max {
			max = v
			key = k
		}
	}
	return key
}