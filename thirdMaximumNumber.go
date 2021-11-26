func thirdMax(nums []int) int {
    uniqArray := uniq(nums)
    
    if len(uniqArray) < 3 {
        max := 0
        for _, v := range nums {
            if v > max { max = v }
        }
        return max
    }
    
    sortedArray := sort(uniqArray)
    
    return sortedArray[len(sortedArray)-3]
}

func uniq(target []int) (uniq []int) {
    m := map[int]bool{}
    
    for _, v := range target {
        if !m[v] {
            m[v] = true
            uniq = append(uniq, v)
        }
    }
    
    return
}

func sort(val []int) (res []int) {
    if len(val) < 2 {
        return val
    }
    
    pivot := val[0]
    
    left := []int{}
    right := []int{}
    for _, v := range val[1:] {
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
