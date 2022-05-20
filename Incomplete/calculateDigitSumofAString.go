func digitSum(s string, k int) string {
   for len(s) > k {
        n, newS := 0, ""
        for i := 0; i < len(s); i++ {
            n += int(s[i] - '0')
            if (i + 1) % k == 0 || i == len(s) - 1 {
                tmp := strconv.Itoa(n)
                newS += tmp
                n = 0
            }
        }
        s = newS
    }   
    return s
}
