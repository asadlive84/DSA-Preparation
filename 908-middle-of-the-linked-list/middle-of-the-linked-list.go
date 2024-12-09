/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	if head == nil {
		return &ListNode{}
	}

	curr := head

	list := []int{}
	count := 0
	for curr != nil {
		list = append(list, curr.Val)
		curr = curr.Next
		count++
	}
	mid := (count / 2) + 1
	//fmt.Println("::list::", list, " count:: ", count, " mid::", mid)
	curr = head
	for curr != nil {
		//fmt.Println("count===>", mid)
        mid--
		if mid == 0 {
			break
		}
		curr = curr.Next
		
	}
	//fmt.Println("::middle::", curr)
	return curr
}