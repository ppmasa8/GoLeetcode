func smallestRangeI(nums []int, k int) int {
    min, max := min(nums), max(nums)
    if max - min - (2 * k) > 0 {
        return max - min - (2 * k)
    }
    return 0
}

func max(nums []int) int {
    max := 0
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}

func min(nums []int) int {
    min := 1000000
    for _, v := range nums {
        if min > v {
            min = v
        }
    }
    return min
}
