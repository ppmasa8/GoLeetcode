/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func binaryTreePaths(root *TreeNode) []string {
    res := []string{}
    if root == nil {
        return res
    }
    helper(root, strconv.Itoa(root.Val), &res)
    return res
}

func helper(root *TreeNode, str string, res *[]string) {
    if root.Left != nil {
        helper(root.Left, str + "->" + strconv.Itoa(root.Left.Val), res)
    }
    if root.Right != nil {
        helper(root.Right, str + "->" + strconv.Itoa(root.Right.Val), res)
    }
    if root.Left == nil && root.Right == nil {
        *res = append(*res, str)
        return
    }
}