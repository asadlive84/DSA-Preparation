func containsNearbyDuplicate(nums []int, k int) bool {
    window := make(map[int]bool) // A map to simulate the sliding window (acting as a set)

    for i, value := range nums {
        // If the value already exists in the window, return true
        if window[value] {
            return true
        }

        // Add the current value to the window
        window[value] = true

        // Ensure the window size does not exceed k
        if len(window) > k {
            // Remove the oldest element from the window
            delete(window, nums[i-k])
        }
    }

    return false // No duplicates found within range k
}
