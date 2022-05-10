/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    res := 0

	postOrder(root, &res)
	return res
}

func postOrder(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    
    left, right := postOrder(root.Left, res), postOrder(root.Right, res)
    
    *res += abs(left - right)
    
    return left + right + root.Val
}

func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}
