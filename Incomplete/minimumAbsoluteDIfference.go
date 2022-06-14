func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    res:=[][]int{}
    min:=999
    for i:=1;i<len(arr);i++ {
        dx:=arr[i] - arr[i-1]
        if min>dx {
            min = dx
            res = [][]int{}
        }
        if min == dx {
            res = append(res, []int{arr[i-1], arr[i]})
        }
        
    }
    return res
}
