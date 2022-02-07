import "fmt"

func numOfStrings(patterns []string, word string) int {
    ans, n, tmp := 0, len(word), make(map[string]bool)
    for i := 1; i <= n; i++ {
        for j := 0; j + i <= n; j++ {
            tmp[word[j:j+i]] = true
        }
    }
    fmt.Println(tmp)
    for _, p := range patterns { if tmp[p] { ans++ } }
    return ans
}
