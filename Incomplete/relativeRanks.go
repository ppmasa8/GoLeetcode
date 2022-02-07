type Athlet struct {
    Scores   int
    Position int
}

func findRelativeRanks(nums []int) []string {
    athlets := make([]Athlet, len(nums))
    ans     := make([]string, len(nums))
    
    for i := 0; i < len(nums); i++ {
        scores := nums[i]
        athlets[i].Scores   = scores
        athlets[i].Position = i
    }
    
    sort.Slice(athlets, func(i, j int) bool {
        return athlets[i].Scores > athlets[j].Scores
    })
    
    for i := 0; i < len(athlets); i++ {
        athlet := athlets[i]
        
        if i == 0 {
            ans[athlet.Position] = "Gold Medal"
        } else if i == 1 {
            ans[athlet.Position] = "Silver Medal"
        } else if i == 2 {
            ans[athlet.Position] = "Bronze Medal"
        } else {
            ans[athlet.Position] = strconv.Itoa(i + 1)
        }
    }
    
    return ans
}
