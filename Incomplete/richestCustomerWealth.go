func maximumWealth(accounts [][]int) int {
    wealth := []int{}
    for _, arr := range accounts {
        wealth = append(wealth, sum(arr))
    }
    return max(wealth)
}

func sum(arr []int) int {
    sum := 0
    for _, v := range arr {
        sum += v
    }
    return sum
}

func max(arr []int) int {
    max := 0
    for _, v := range arr {
        if max < v {
            max = v
        }
    }
    return max
}
