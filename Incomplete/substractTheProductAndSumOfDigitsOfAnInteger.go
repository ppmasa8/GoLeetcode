func subtractProductAndSum(n int) int {
    sum:=0
    pro:=1
    for n > 0 {
        sum += n%10
        pro *= n%10
        n /= 10
    }
    return pro - sum
}
