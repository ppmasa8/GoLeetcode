func largestGoodInteger(num string) string {
    arr := []string{}
    tmp := ""
    for _, v := range num {
        if tmp == "" {
            tmp = tmp + string(v)
        } else if string(tmp[0]) == string(v) && len(tmp) < 3 {
            tmp = tmp + string(v)
        } else if len(tmp) == 3 {
            arr = append(arr, tmp)
            tmp = string(v)
        } else {
            tmp = string(v)
        }
    }
    
    if len(tmp) == 3 {
        arr = append(arr, tmp)
    }
    
    // When arr is nil.
    if len(arr) == 0 {
        return ""
    }
    
    largeInt := 0
    for _, v := range arr {
        i, _ := strconv.Atoi(v)
        if largeInt < i {
            largeInt = i
        }
    }
    
    str := strconv.Itoa(largeInt)
    if str == "0" {
        str = "000"
    }
    return str
}
