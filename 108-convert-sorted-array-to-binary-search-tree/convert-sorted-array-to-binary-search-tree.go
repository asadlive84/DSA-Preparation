/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

	if nums == nil || len(nums) == 0 {
		return nil
	}

	mid := (len(nums) - 1) / 2

	root := &TreeNode{Val: nums[mid]}

	stack := []struct {
		Root   *TreeNode
		Left   int
		Right  int
		IsLeft bool
	}{
		{root, 0, mid - 1, true},
		{root, mid + 1, len(nums) - 1, false},
	}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left, right := curr.Left, curr.Right

		if left > right {
			continue
		}

		mid = (left + right) / 2

		if mid >= len(nums) {
			continue
		}

		newNode := &TreeNode{Val: nums[mid]}

		if curr.IsLeft {
			curr.Root.Left = newNode
		} else {
			curr.Root.Right = newNode
		}

		stack = append(stack, struct {
			Root   *TreeNode
			Left   int
			Right  int
			IsLeft bool
		}{newNode, left, mid - 1, true})

		stack = append(stack, struct {
			Root   *TreeNode
			Left   int
			Right  int
			IsLeft bool
		}{newNode, mid + 1, right, false})
	}

	return root
}