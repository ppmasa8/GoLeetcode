func replaceDigits(s string) string {
    ans := []byte(s)
    for i := 1; i < len(s); i += 2 {
        ans[i] = ans[i - 1] + ans[i] - '0'
    }
    return string(ans)
}
