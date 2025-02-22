func twoSum(nums []int, target int) []int {
    m := make(map[int]int) // Map to store value and index

    for i, num := range nums {
        complement := target - num
        if j, found := m[complement]; found {
            return []int{j, i} // Return indices of the two numbers
        }
        m[num] = i // Store current number with its index
    }

    return []int{-1, -1} // If no solution is found
}
