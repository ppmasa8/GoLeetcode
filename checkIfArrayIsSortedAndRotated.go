import (
    "reflect"
)

func check(nums []int) bool {
    sorted := sort(nums)
    if reflect.DeepEqual(nums, sorted) { return true }
    for i := 1; i < len(nums); i++ {
        sorted = append(sorted[1:], sorted[0])
        if reflect.DeepEqual(nums, sorted) { return true }
    }
    return false
}

func sort(arr []int) (res []int) {
    if len(arr) < 2 { return arr }
    
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
    
    return res
}
