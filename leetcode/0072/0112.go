func hasPathSum(root *TreeNode, targetSum int) bool {
    var dfs func(*TreeNode,int)
    ans := false
    dfs = func(n *TreeNode,s int){
        if n == nil{
            return
        }
        s +=  n.Val
        
        if n.Left == nil && n.Right== nil{
            if s == targetSum{
                ans = true
            }
        }
        dfs(n.Left,s)
        dfs(n.Right,s)
        
    }
    dfs(root, 0)
    return ans
}