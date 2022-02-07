import "fmt"
import "strconv"

func freqAlphabets(s string) string {
   list := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
    arr := []string{}
    res := ""
    for i := 0; i < len(s); i++ {
        if string(s[i]) == "#" {
            arr = arr[0:len(arr) - 2]
            arr = append(arr, string(s[i-2]) + string(s[i-1]))
        } else {
            arr = append(arr, string(s[i]))
        }
    }
    
    for _, v := range arr {
        num, _ := strconv.Atoi(v)
        res += list[num - 1]
    }
    return res
}
