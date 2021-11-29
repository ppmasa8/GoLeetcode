func reverseVowels(s string) string {
    stackVowels := []string{}
    res := []string{}
    
    for i := 0; i < len(s); i++ {
        if detectVowels(string(s[i])) {
            stackVowels = append(stackVowels, string(s[i]))
        }
    }
        
    stackVowels = reverseString(stackVowels)
    
    for i := 0; i < len(s); i++ {
        if detectVowels(string(s[i])) {
            res = append(res, stackVowels[0])
            stackVowels = stackVowels[1:]
        } else {
            res = append(res, string(s[i]))
        }
    }
    
    return strings.Join(res, "")
}

func detectVowels(s string) bool {
    vowels := []string{"a", "i", "u", "e", "o", "A", "I", "U", "E", "O"}
    for _, v := range vowels {
        if s == v {
            return true
        }
    }
    return false
}

func reverseString(str []string) []string {
    for i := 0; i < len(str) / 2; i++ {
        str[i], str[len(str) - i - 1] = str[len(str) - i - 1], str[i]
    }
    return str
}
