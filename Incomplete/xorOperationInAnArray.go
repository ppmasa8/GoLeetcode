func xorOperation(n int, start int) int {
    arr:=[]int{}
    xor:=0
 
    for i := 1; i <= n; i++ {
        arr = append(arr, start+2*(i-1))
    }
    for _, v := range arr {
        xor = xor ^ v
    }
    return xor
}
