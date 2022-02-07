func arraySign(nums []int) int {
    cnt := 0
    for _, v := range nums {
        if v == 0 { return 0 }
        v > 0 ? cnt++ : cnt--
    }
    return cnt
}
