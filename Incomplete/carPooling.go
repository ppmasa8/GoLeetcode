import "fmt"

func carPooling(trips [][]int, capacity int) bool {
    tmp, temp := make([]int, 1001), 0
    for _, v := range trips {
        fmt.Println(v[1])
        tmp[v[1]] += v[0]
        tmp[v[2]] -= v[0]
    }
    for _, v := range tmp {
        temp += v
        if temp > capacity { return false }
    }
    return true
}
