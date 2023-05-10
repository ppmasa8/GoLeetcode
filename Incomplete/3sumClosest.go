func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    ans, n := math.MaxInt32, len(nums)
    for i := 0; i < n - 2; i++ {
        low, high := i + 1, n - 1
        for low < high {
            curSum := nums[low] + nums[high] + nums[i]
            if curSum > target {
                high--
            } else {
                low++
            }
            if abs(curSum - target) < abs(ans - target) { ans = curSum }
        }
    }
    return ans
}

func abs(a int) int {
    if a > 0 { return a }
    return -a
}
