var index = map[byte]string {
    '1': "",
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
    '0': " ",
}


func letterCombinations(digits string) []string {
    res := []string{}
    if len(digits) == 0 {
        return res
    }
    
    dfs(digits, 0, "", &res)
    
    return res
}

func dfs(digits string, pos int, s string, strs *[]string) {
    if pos == len(digits) {
        *strs = append(*strs, s)
        // fmt.Println("pos equals digits length")
        return
    }
    
    for _, v := range index[digits[pos]] {
        // fmt.Println(digits, pos, s, strs)
        dfs(digits, pos+1, s+string(v), strs)
    }
    return
}
