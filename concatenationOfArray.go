func getConcatenation(nums []int) []int {
    res := []int{}
    for _, v := range nums {
        res = append(res, v)
    }
    for _, v := range nums {
        res = append(res, v)
    }
    return res
}
