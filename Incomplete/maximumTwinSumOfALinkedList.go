/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    tmp := []int{}
    for head != nil {
        tmp = append(tmp, head.Val)
        head = head.Next
    }
    res := 0
    for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1  {
        if res < tmp[i] + tmp[j] {
            res = tmp[i] + tmp[j]
        }
    }
    return res
}
