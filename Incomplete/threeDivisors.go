func isThree(n int) bool {
    foundOne := false
    for i := 2; i <= n/2; i++ {
        if n % i == 0 {
            if foundOne { return false }
            foundOne = true
        }
    }
    return foundOne
}
