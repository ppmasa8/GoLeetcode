func findGCD(nums []int) int {
    arr := findMaxMin(nums)
    max, min := arr[1], arr[0]
    for i := min; i > 0; i-- {
        if max % i == 0 {
            if min % i == 0 {
                return i
            }
        }
    }
    return 1
}

func findMaxMin(nums []int) (arr []int) {
    max := 0
    min := 1000
    
    for _, v := range nums {
        if v > max { max = v }
        if v < min { min = v }
    }
    
    return []int{min, max}
}
