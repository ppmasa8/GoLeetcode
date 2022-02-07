func diStringMatch(s string) []int {
    i, d := 0, len(s)
    arr := []int{}
    
    for _, v := range s {
        if v == 'I' {
            arr = append(arr, i)
            i++
        } else if v == 'D' {
            arr = append(arr, d)
            d--
        }
    }
    arr = append(arr, i)
    return arr
}
