func convertBST(root *TreeNode) *TreeNode {
    s := 0
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode){
        if node == nil{
            return
        }
        dfs(node.Right)
        s+=node.Val
        node.Val = s
        dfs(node.Left)
        return
    }
    dfs(root)
    return root
}