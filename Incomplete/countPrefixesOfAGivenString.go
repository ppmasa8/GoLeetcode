func countPrefixes(words []string, s string) int {
    length := len(s)
    strings := ""
    count := 0
    for i := 0; i <= length; i++ {
        strings = s[:i]
        for _, v := range words {
            if v == strings {
                count++
            }
        }
    }
    return count
}
