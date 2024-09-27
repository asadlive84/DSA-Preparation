package main

import "fmt"

func mergeSortedArray(a []int, b []int) []int {

	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	newArr := make([]int, 0, len(a)+len(b))

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			newArr = append(newArr, a[i])
			i++
		} else {
			newArr = append(newArr, b[j])
			j++
		}
	}

	return newArr

}

func main() {

	a := []int{0, 3, 4, 31}
	b := []int{4, 6, 30}

	fmt.Println("========================================")
	fmt.Printf("mergeSortedArray %+v\n", mergeSortedArray(a, b))
	fmt.Println("========================================")

}
