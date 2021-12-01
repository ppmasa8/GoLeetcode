func fib(n int) int {
    list := []int{}
    for i := 0; i <= n; i++ {
        list = recursionFormula(i, list)
    }
    return list[len(list) - 1]
}

func recursionFormula(n int, arr []int) []int {
    if n == 0 { 
        arr = append(arr, 0)
        return arr
    }
    
    if n == 1 { 
        arr = append(arr, 1)
        return arr
    }
    
    num := arr[len(arr) - 1] + arr[len(arr) - 2]
    arr = append(arr, num)
    return arr
}
