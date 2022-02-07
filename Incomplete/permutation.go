func permute(nums []int) [][]int {
    return permutation(nums)
}

func remove (arr []int, i int) []int {
    tmp := make([]int, len(arr))
    copy(tmp, arr)
    return append(tmp[0:i], tmp[i+1:]...)
}

func permutation(arr []int) (res [][]int) {
    if len(arr) == 1 { return append(res, arr) }
    for i, a := range arr {
        for _, b := range permutation(remove(arr, i)) {
            res = append(res, append([]int{a}, b...))
        }
    }
    return res
}
