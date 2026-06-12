func maxAncestorDiff(root *TreeNode) (ans int) {
    var dfs func(*TreeNode, int, int)
    dfs = func(node *TreeNode,mi,ma int){
        if node == nil{
            return
        }
        m1 := absInt(ma-node.Val)
        m2 := absInt(node.Val-mi)
        miaa := min(node.Val,mi)
        maaa := max(node.Val,ma)
        ans = max(ans,m1,m2)
        dfs(node.Left,miaa,maaa)
        dfs(node.Right,miaa,maaa)
    }
    dfs(root,root.Val,root.Val)
    return 
}
func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}