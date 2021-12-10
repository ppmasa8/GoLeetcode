func maxDistance(colors []int) int {
    ans, n := 0, len(colors)
    for i := 1; i < n; i++ {
        if colors[0] != colors[i] {
            ans = max(ans, i - 0)
        }
    }
    for i := n - 2; i >= 0; i-- {
        if colors[n - 1] != colors[i] {
            ans = max(ans, n - 1 - i)
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
