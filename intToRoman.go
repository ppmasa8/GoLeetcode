func intToRoman(num int) string {
    var str strings.Builder
    sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    
    for i := 0; i < 13; i++ {
        j := num / val[i]
        num %= val[i]
        for j > 0 {
            str.WriteString(sym[i])
            j--
        }
    }
    return str.String()
}
