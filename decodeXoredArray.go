func decode(encoded []int, first int) []int {
    res := []int{}
    res = append(res, first)
    for _, v := range encoded {
        res = append(res, res[len(res) - 1] ^ v)
    }
    return res
}
