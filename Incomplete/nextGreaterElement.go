func nextGreaterElement(nums1 []int, nums2 []int) []int {
  for i, v := range nums1 {
		imIn := false
		res := -1
		for _, p := range nums2 {
			if v == p {
				imIn = true
			}
			if imIn {
				if v < p {
					res = p
					break
				}
			}
		}
		nums1[i] = res
	}
    return nums1
}
