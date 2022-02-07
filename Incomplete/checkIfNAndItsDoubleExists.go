import "fmt"
func checkIfExist(arr []int) bool {
    for _, v := range arr {
        if exist(v * 2, arr) {
            return true
        }
    }
    return false
}

func exist(num int, arr []int) bool {
    fmt.Println(zeroCnt(arr))
    if zeroCnt(arr) >= 2 { return true }
    for _, v := range arr {
        if v == 0 { continue }
        if num == v {
            return true
        }
    }
    return false
}

func zeroCnt(arr []int) (cnt int) {
    for _, v := range arr {
        if v == 0 { cnt++ }
    }
    return cnt
}
