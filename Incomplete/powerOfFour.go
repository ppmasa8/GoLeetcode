func isPowerOfFour(n int) bool {
    if n == 0 {
        return false
    }
    a := math.Log(float64(n)) / math.Log(4)
    return math.Floor(a) == a
}
