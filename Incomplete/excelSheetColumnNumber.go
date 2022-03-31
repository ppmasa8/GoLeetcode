func titleToNumber(columnTitle string) int {
    n := 0
	for i := 0; i < len(columnTitle); i++ {

		n = (int(columnTitle[i]) - 64) + 26*n
	}
	return n
}
