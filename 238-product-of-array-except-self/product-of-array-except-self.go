func productExceptSelf(arr []int) []int {
	arr1, arr2 := make([]int, len(arr), len(arr)), make([]int, len(arr), len(arr))
	prev, last := 1, 1
	arr1[0], arr2[len(arr2)-1] = prev, last
	for i := 1; i < len(arr); i++ {
		prev = prev * arr[i-1]
		arr1[i] = prev
	}

	for j := len(arr) - 2; j >= 0; j-- {
		last = last * arr[j+1]
		arr2[j] = last
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = arr1[i] * arr2[i]
	}
	fmt.Println(arr1)
	return arr
}