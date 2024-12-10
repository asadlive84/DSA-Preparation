func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	k := m + n - 1

	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] < nums2[p2] {
			nums1[k] = nums2[p2]
			p2--
			k--

		} else if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[k] = nums1[p1]
			p1--
			k--
		} else {
			nums1[k] = nums2[p2]
			p2--
			k--
		}
	}

}