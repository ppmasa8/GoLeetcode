func divideArray(nums []int) bool {
    hash := map[int]int{}
    for _, v := range nums {
        hash[v]++
    }
    
    for _, v := range hash {
        if v % 2 != 0 {
            return false
        }
    }
    return true
}
