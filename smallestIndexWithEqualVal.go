func smallestEqual(nums []int) int {
    for i, v := range nums {
        if v % 10 == i % 10 {
            return i
        }
    }
    return -1
}
