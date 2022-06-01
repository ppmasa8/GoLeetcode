ans := []int{}
    for i, j := 0, n; i < n && j < 2 * n; i, j = i + 1, j + 1 { ans = append(ans, []int{nums[i], nums[j]}...) }
    return ans
