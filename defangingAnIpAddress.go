func defangIPaddr(address string) string {
    res := []string{}
    for i := 0; i < len(address); i++ {
        if string(address[i]) == "." {
            res = append(res, "[")
            res = append(res, string(address[i]))
            res = append(res, "]")
        } else {
            res = append(res, string(address[i]))
        }
    }
    return strings.Join(res, "")
}
