func balancedStringSplit(s string) int {
    balance, balancedStrings := 0, 0
    for i, r := range s {
        if string(r) == "L" {
            balance++
        }
        
        if string(r) == "R" {
            balance--
        }
        
        if i != 0 && balance == 0 {
            balancedStrings++
        }
    }
    
    return balancedStrings
}
