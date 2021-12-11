func reversePrefix(word string, ch byte) string {
    var ans string
    
    for i := 0; i < len(word); i++ {
        if word[i] == ch {
            pivot := i
            
            for i >= 0 {
                ans += string(word[i])
                i--
            }
            
            return ans + word[pivot + 1:]
        }
    }
    
    return word
}
