func judgeCircle(moves string) bool {
    r,l,u,d:=0,0,0,0
    for _, v := range moves {
        if string(v) == "R" {
            r++
        } else if string(v) == "L" {
            l++
        } else if string(v) == "U" {
            u++
        } else {
            d++
        }
    }
    
    if r - l == 0 && u - d == 0 {
        return true
    }
    return false
}
