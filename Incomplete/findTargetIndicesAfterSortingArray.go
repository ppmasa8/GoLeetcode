func targetIndices(nums []int, target int) []int {
    sortedNums := sort(nums)
    res := []int{}
    for i, v := range sortedNums {
        if v == target {
            res = append(res, i)
        } else if v > target {
            return res
        }
    }
    return res
}

func sort(arr []int) (res []int) {
    if len(arr) < 2 {
        return arr
    }
    
    pivot := arr[0]
    
    left, right := []int{}, []int{}
    for _, v := range arr[1:] {
        if v > pivot {
            right = append(right, v)
        } else {
            left = append(left, v)
        }
    }
    
    left = sort(left)
    right = sort(right)
    
    res = append(left, pivot)
    res = append(res, right...)
    return
}
