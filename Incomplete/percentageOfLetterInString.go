func percentageLetter(s string, letter byte) int {
    cnt, lenS := 0, len(s)
    for _, v := range s {
        fmt.Println(string(v))
        if string(v) == string(letter) {
            cnt++
        }
    }
    return cnt * 100 / lenS 
}
