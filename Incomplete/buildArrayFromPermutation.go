func buildArray(nums []int) []int {
    arr := []int{}
    for _, v := range nums {
        arr = append(arr, nums[v])
    }
    return arr
}
