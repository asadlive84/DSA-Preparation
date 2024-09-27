package main

import "fmt"

func twoSum(nums []int, target int) []int {

	position := make([]int, 0)

	for index, num := range nums {
		targetNextSum := target - num
		for j := index + 1; j < len(nums); j++ {
			if targetNextSum == nums[j] {
				val := j
				if index == j {
					val++
				}
				position = append(position, index)
				position = append(position, val)
				return position
			}
		}

	}

	return position
}

func main() {

	// nums := []int{2, 7, 11, 15}
	// nums := []int{3, 2, 4}
	// nums := []int{3, 2, 3}
	// nums := []int{2, 5, 5, 11}
	nums := []int{1, 3, 4, 2}

	fmt.Println("========================================")
	fmt.Printf("%+v\n", twoSum(nums, 6))
	fmt.Println("========================================")
}
