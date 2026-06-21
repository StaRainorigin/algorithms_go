func pathSum(root *TreeNode, targetSum int) ( ans [][]int )  {
    path := []int{}
    var dfs func(*TreeNode,int)
    dfs = func(node *TreeNode,sum int){
        if node== nil {
            return
        }
        path = append(path, node.Val)
        curSum := sum + node.Val
        if node.Left == nil && node.Right ==nil && curSum == targetSum{
            ans = append(ans, slices.Clone(path))
        }else{
            dfs(node.Left,curSum)
            dfs(node.Right,curSum)
        }
        
        path = path[:len(path)-1]
    }
    dfs(root,0)
    return
}