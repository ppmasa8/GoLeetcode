import "strconv"

func maximum69Number (num int) int {
    str := strconv.Itoa(num)
    rs := []rune(str)
    for i := 0; i < len(str); i++ {
        if string(rs[i]) == "6" {
            rs[i] = 57 //　It's an int32 rune value assigned to "9".
            res, _ := strconv.Atoi(string(rs))
            return res
        }
    }
    res, _ := strconv.Atoi(string(rs))
    return res
}
