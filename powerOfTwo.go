import "strconv"

func isPowerOfTwo(n int) bool {
    if n < 0 { return false }
    cntOne := 0
    str := strconv.FormatInt(int64(n), 2)
    for _, v := range str {
        if string(v) == "1" { cntOne++ }
    }
    return cntOne == 1
}
