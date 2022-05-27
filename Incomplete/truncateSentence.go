func truncateSentence(s string, k int) string {
    res := []string{}
    spaceCnt := 0
    for _, v := range s {
        if string(v) == " " {
            spaceCnt++
        }
        if spaceCnt >= k {
            return strings.Join(res, "")
        }
        res = append(res, string(v))
        fmt.Println(res)
    }
    return strings.Join(res, "")
}
