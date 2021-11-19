func addDigits(num int) int {
    for {
        if num < 10 {
            return num
        } else {
            num = culcDigits(num)
        }
    }
    
    return num
}

func culcDigits(num int) (cnt int) {
    str := strconv.Itoa(num)
    slice := strings.Split(str, "")
    for i := 0; i < len(slice); i++ {
        n, _ := strconv.Atoi(slice[i])
        cnt += n
    }
    return cnt
}
