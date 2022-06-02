/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    return recursion(root, low, high, 0)
}

func recursion(root *TreeNode, low int, high int, sum int) int {
    if root == nil {
        return sum
    }
    
    if root.Val >= low && root.Val <= high {
        sum += root.Val
    }
    
    sum = recursion(root.Left, low, high, sum)
    sum = recursion(root.Right, low, high, sum)
    
    return sum
}
