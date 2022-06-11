func squareIsWhite(coordinates string) bool {
    odd, even := 0, 0
    for _, v := range coordinates {
        if v % 2 == 0 {
            even++
        } else {
            odd++
        }
    }
    
    if odd == 2 || even == 2 {
        return false
    }
    return true
}
