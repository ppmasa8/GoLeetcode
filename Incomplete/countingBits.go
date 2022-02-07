func countBits(n int) []int {
	arr := []int{}
	cntOne := 0
    for i := 0; i <= n; i++ {
		str := strconv.FormatInt(int64(i), 2)
		for _, v := range str {
			if string(v) == "1" { cntOne++ }
		}
		arr = append(arr, cntOne)
		cntOne = 0
	}
	return arr
}