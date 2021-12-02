import "fmt"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    sumA, sumB := sum(aliceSizes), sum(bobSizes)
    
    for _, x := range aliceSizes {
        if exist(x + (sumB - sumA) / 2, bobSizes) {
            return []int{x, x + (sumB - sumA) / 2 }
        }
    }
    
    // It's impossible.
    return []int{}
}

func sum(arr []int) (sum int) {
    for _, v := range arr {
        sum += v
    }
    return sum
}

func exist(num int, arr []int) bool {
    for _, v := range arr {
        if num == v {
            return true
        }
    }
    return false
}
