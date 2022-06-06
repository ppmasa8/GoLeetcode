/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    list, tmp := &ListNode{}, list1
    list.Next = list1
    for a > 0 || b > 0 {
        a--
        b--
        if a > 0 {
            list1 = list1.Next
        }
        if b >= 0 {
            tmp = tmp.Next
        }
    }
    list1.Next = list2
    for list2.Next != nil {
        list2 = list2.Next
    }
    list2.Next = tmp.Next
    return list.Next
}
