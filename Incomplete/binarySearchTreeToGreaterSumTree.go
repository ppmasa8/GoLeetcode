/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    helper(root, 0)
    return root
}

func helper(root *TreeNode, n int) int {
    if root.Left == nil && root.Right == nil {
        root.Val += n
        return root.Val - n
    }
    res := root.Val
    if root.Right != nil {
        res += helper(root.Right, n)
    }
    root.Val = res + n
    n = root.Val
    if root.Left != nil {
        res += helper(root.Left, n)
    }
    return res
}
