func maxPower(s string) int {
    if len(s) == 1 { return 1 }
    cnt := 1
    var arr []int
    
    for i := 0; i < len(s) - 1; i++ {
        if s[i] == s[i+1] {
            cnt++
        } else {
            arr = append(arr, cnt)
            cnt = 1
        }
    }
    arr = append(arr, cnt)
    return max(arr)
}

func max(nums []int) int {
    max := 0
    for i := 0; i < len(nums); i++ {
        if max < nums[i] {
            max = nums[i]
        }
    }
    return max
}
