func evaluateTree(root *TreeNode) bool {
    if root.Val == 0 || root.Val ==1{
        return root.Val == 1
    }
    if root.Val == 2{
        return evaluateTree(root.Left) || evaluateTree(root.Right)
    }
    if root.Val == 3{
        return evaluateTree(root.Left) && evaluateTree(root.Right)
    }
    return false
}