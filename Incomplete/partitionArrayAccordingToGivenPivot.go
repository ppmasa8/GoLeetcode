func pivotArray(nums []int, pivot int) []int {
    left, right := []int{}, []int{}
    lenNums := len(nums)
    for _, v := range nums {
        if v > pivot {
            right = append(right, v)
        } else if v < pivot {
            left = append(left, v)
        }
    }
    
    for i := len(left)+len(right); i < lenNums; i++ {
        left = append(left, pivot)
    }
    
    for _, v := range right {
        left = append(left, v)
    }
    return left
}
