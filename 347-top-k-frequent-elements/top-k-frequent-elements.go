func topKFrequent(arr []int, k int) []int {

	if len(arr) == 0 {
		return []int{}
	}

	elementsMap := make(map[int]int)

	for _, data := range arr {
		elementsMap[data]++
	}

	newArr := make([]int, 0, len(elementsMap))

	for data := range elementsMap {
		newArr = append(newArr, data)
	}

	sort.Slice(newArr, func(a, b int) bool {
		return elementsMap[newArr[a]] > elementsMap[newArr[b]]
	})

	return newArr[:k]

}