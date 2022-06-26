func sortEvenOdd(nums []int) []int {
    s1 := make([]int, 0)
    s2 := make([]int, 0)
    for i, num := range nums {
        if i % 2 == 1 {
            s1 = append(s1, num)
        } else {
            s2 = append(s2, num)
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(s1)))
    sort.Ints(s2)
    p1 := 0
    p2 := 0
    for i := 0 ; i < len(nums); i++ {
        if i % 2 == 1 {
            nums[i] = s1[p1]
            p1++
        } else {
            nums[i] = s2[p2]
            p2++
        }
    }
    return nums
}
