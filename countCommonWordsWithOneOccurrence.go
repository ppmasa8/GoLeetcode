func countWords(words1 []string, words2 []string) int {
    words := []string{}
    cnt := 0
    if len(words1) > len(words2) {
        words = words1
    } else {
        words = words2
    }
    
    for _, v := range words {
        if count(v, words1) == 1 {
            if count(v, words2) == 1 {
                cnt++
            }
        }
    }
    return cnt
}

func count(s string, arr []string) (cnt int) {
    for _, v := range arr {
        if s == v {
            cnt++
        }
    }
    return
}
