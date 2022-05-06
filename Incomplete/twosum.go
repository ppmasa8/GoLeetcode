func twoSum(nums []int, target int) []int {
  m := make(map[int]int, len(nums))

  for i, num := range nums (
    if j, v := m[target - num]; ok {
      return []int{j, i}
    }
    m[num] = i
  }
  return []int{0, 0}
}
