func findTilt(root *TreeNode) int {
    ans := 0
    var dfs func(*TreeNode)int
    dfs = func(node *TreeNode)int{
        if node == nil{
            return 0
        }
        l := dfs(node.Left)
        r := dfs(node.Right)
        ans = ans + abs(l-r)
        return node.Val + l +r
    }
    dfs(root)  
    return ans
}
func abs(i int)int{
    if i < 0{
        return -i
    }
    return i
}