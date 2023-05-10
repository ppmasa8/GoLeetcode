func findNumbers(nums []int) int {
    evenCnt := 0
    for _, v := range nums {
        str := strconv.Itoa(v)
        if len(str) % 2 == 0 {
            evenCnt += 1
        }
    }
    
    return evenCnt
}
