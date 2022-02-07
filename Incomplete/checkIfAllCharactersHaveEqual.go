func areOccurrencesEqual(s string) bool {
    hash := map[rune]int{}
    
    for _, v := range s {
        hash[v]++
    }
    
    ref := hash[rune(s[0])]

    for _, v := range hash {
        if v != ref {
            return false
        }
    }
    
    return true
}
