func restoreString(s string, indices []int) string {
    num := len(indices)+1
    res := make([]string, num)
    for k, v := range indices {
        res[v] = string(s[k])   
    }
    return strings.Join(res, "")
}
