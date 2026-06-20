func amountOfTime(root *TreeNode, start int) (ans int) {
    a:=0
    var dfs func(*TreeNode) (int, bool)
    dfs = func(node *TreeNode) (int, bool) {
        if node == nil{
            return 0,false
        }
        l,la:= dfs(node.Left)
        r,ra:=dfs(node.Right)
        if node.Val == start{
            a = max(l,r) 
            if l>r{
                la = true
            }
            ra = true
        }
        if la || ra{
            ans = max(ans,l+r)
            if la{
                return l+1 ,true
            }
            return r+1,true
        }
        return max(l,r)+1 , la || ra
    }
    dfs(root)
    return max(ans-a,a)
}