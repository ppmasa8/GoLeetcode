func mostWordsFound(sentences []string) int {
    ans := 0
    for _, sentence := range sentences {
        max := helper(sentence)
        if max > ans { ans = max }
    }
    return ans
}

func helper(sentence string) int {
    n := 0
    for _, c := range sentence {
        if c == ' ' { n++ }
    }
    return n + 1
}
