func runningSum(nums []int) []int {
    tmp := 0
    arr := []int{}
    for _, v := range nums {
        tmp += v
        arr = append(arr, tmp)
    }
    return arr
}
