/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var res int
    dfs(root, &res)
    return res
}

func dfs(root *TreeNode, res *int) (left int, right int) {
    if root == nil {
        return -1, -1
    }
    
    a, b := dfs(root.Left, res)
    c, d := dfs(root.Right, res)
    left = max(a,b) + 1
    right = max(c,d) + 1
    *res = max(*res, left + right)
    return left, right
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
