func climbStairs(n int) int {
    if n < 3 { return n }
    
    one, two := 1, 1
    for i := 0; i < n-2; i++ {
        next := one + two
        one = two
        two = next
    }
    return one + two
}
