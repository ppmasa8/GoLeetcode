func minimumMoves(s string) int {
    res := 0
    for i := 0; i < len(s); {
        if string(s[i]) == "O" {
            i++
        } else {
            i+=3
            res++
        }
    }
    return res
}
