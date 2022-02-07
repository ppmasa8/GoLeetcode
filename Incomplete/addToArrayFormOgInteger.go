func addToArrayForm(num []int, k int) []int {
    res := []int{}
    i := len(num) - 1
    for i >= 0 || k > 0 {
        if i >= 0 { k += num[i] }
        res = append(res, k%10)
        k /= 10
        i--
    }
    for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
