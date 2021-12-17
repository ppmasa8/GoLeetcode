func maxProductDifference(nums []int) int {
    sortedArr := sort(nums)
    return sortedArr[len(sortedArr)-1]*sortedArr[len(sortedArr)-2] - sortedArr[0]*sortedArr[1]
}

func sort(nums []int) []int {
    var res []int
    if len(nums) < 2 { return nums }
    
    mid := int(len(nums) / 2)
    r := sort(nums[:mid])
    l := sort(nums[mid:])
    i, j := 0, 0
    
    for i < len(r) && j < len(l) {
        if r[i] > l[j] {
            res = append(res, l[j])
            j++
        } else {
            res = append(res, r[i])
            i++
        }
    }
    
    res = append(res, r[i:]...)
    res = append(res, l[j:]...)
    
    return res
}
