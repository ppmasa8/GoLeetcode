func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 { return nil }
    index := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[index] { index = i }
    }
    return &TreeNode{
        Val: nums[index],
        Left: constructMaximumBinaryTree(nums[:index]),
        Right: constructMaximumBinaryTree(nums[index+1:]),
    }
}
