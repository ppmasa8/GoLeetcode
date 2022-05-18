func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    key, cnt := 0, 0
    if ruleKey == "color" {
        key = 1
    } else if ruleKey == "name" {
        key = 2
    }
    
    for _, v := range items {
        if v[key] == ruleValue {
            cnt++
        }
    }
    return cnt
}
