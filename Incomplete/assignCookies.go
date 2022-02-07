import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    firstLen := len(g)
    for _, v := range s {
        if len(g) == 0 {
            break
        }
        if v >= g[0] {
            g = g[1:]
        }
    }
    return firstLen - len(g)
}
