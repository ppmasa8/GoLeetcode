func isMonotonic(a []int) bool {
	countA := 0
	countB := 0
	if len(a) == 1 {
		return true
	}
    
	for i := 1; i < len(a); i++ {
		if a[i-1] <= a[i] {
			countA++
		}
		if a[i-1] >= a[i] {
			countB++
		}
		if countA == len(a)-1 || countB == len(a)-1 {
			return true
		}
	}
    
	return false
}
