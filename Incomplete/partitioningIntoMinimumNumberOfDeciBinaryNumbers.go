func minPartitions(n string) int {
    max := 0
    for len(n) > 0 {
        digit := int(n[0]-'0')
        if digit > max { max = digit }
        n = n[1:]
    }
    return max
}
