func sequentialDigits(low int, high int) []int {
    d, ans := "123456789", []int{}
    l, h := len(strconv.Itoa(low)), len(strconv.Itoa(high))
    for i := l; i <= h; i++ {
        for j := 0; j < 10 - i; j++ {
            n, _ := strconv.Atoi(d[j:j + i])
            if n >= low && n <= high { ans = append(ans, n) }
        }
    }
    return ans
}
